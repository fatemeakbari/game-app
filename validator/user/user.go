package user

import (
	entity "gameapp/entity/userentity"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateRegister(request entity.RegisterRequest) error {

	PhoneNumberPattern := regexp.MustCompile(`^09\d{9}`)

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&request.PhoneNumber, validation.Required, validation.Match(PhoneNumberPattern)),
		validation.Field(&request.Password, validation.Required, validation.Length(5, 50)),
	)

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithMessage(err.Error()).
			WithCodeStatus(errorcodestatus.InvalidProcess).
			WithOperation("ValidateRegister")
	}

	unique, err := v.userRepository.IsPhoneNumberUnique(request.PhoneNumber)

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithCodeStatus(errorcodestatus.InternalError).
			WithOperation("ValidateRegister").
			WithMessage(errormessage.InternalError)
	}

	if !unique {
		return errorhandler.New().
			WithCodeStatus(errorcodestatus.InvalidProcess).
			WithOperation("ValidateRegister").
			WithMessage(errormessage.PhoneNumberDuplicated)
	}

	return nil

}
