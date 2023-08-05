package matchingvalidator

import (
	"fmt"
	entity "gameapp/entity/matching"
	"gameapp/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateAddUserToWaitingListRequest(req entity.AddUserToWaitingListRequest) error {

	err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required, validation.By(v.validateCategory)))

	if err != nil {
		return err
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
