package user

import (
	"errors"
	entity "gameapp/entity/user"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateRegisterRequest(request entity.RegisterRequest) error {

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&request.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberPattern)),
			validation.By(v.checkPhoneNumberIsUnique)),
		validation.Field(&request.Password, validation.Required, validation.Length(5, 50)),
	)

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithMessage(err.Error()).
			WithCodeStatus(errorcodestatus.InvalidProcess).
			WithOperation("ValidateRegisterRequest")
	}

	return nil

}

func (v Validator) checkPhoneNumberIsUnique(value interface{}) error {
	exist, err := v.userRepository.IsPhoneNumberExist(value.(string))

	//dont expect occur err, if error occur must fix it
	if err != nil || exist {
		return errors.New(errormessage.PhoneNumberDuplicated)
	}

	return nil
}
