package user

import (
	"errors"
	entity "gameapp/entity/user"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) ValidateLoginRequest(request entity.LoginRequest) error {

	err := validation.ValidateStruct(&request,
		validation.Field(&request.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberPattern)),
			validation.By(v.checkUserExistByPhoneNumber)),
	)

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithMessage(err.Error()).
			WithCodeStatus(errorcodestatus.InvalidProcess).
			WithOperation("ValidateLoginRequest")
	}

	return nil
}

func (v Validator) checkUserExistByPhoneNumber(value interface{}) error {
	exist, err := v.userRepository.IsPhoneNumberExist(value.(string))

	//dont expect occur err, if error occur must fix it
	if err != nil || !exist {
		return errors.New(errormessage.LoginMessage)
	}

	return nil
}
