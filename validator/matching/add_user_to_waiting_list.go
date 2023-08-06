package matchingvalidator

import (
	"fmt"
	entity "gameapp/entity/matching"
	"gameapp/model"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddUserToWaitingListRequest(req entity.AddUserToWaitingListRequest) error {

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required, validation.By(v.validateCategory)))

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithOperation("ValidateAddUserToWaitingListRequest").
			WithCodeStatus(errorcodestatus.InvalidProcess).
			WithMessage(err.Error())
	}

	return nil
}

func (v Validator) validateCategory(value interface{}) error {
	res, ok := value.(model.Category)

	if !ok || !res.IsValid() {
		return fmt.Errorf("category name is not valid")
	}

	return nil
}
