package user

import (
	"gameapp/model"
)

const (
	phoneNumberPattern = `^09\d{9}`
)

type Repository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	FindUserByPhoneNumber(phoneNumber string) (model.User, error)
}

type Validator struct {
	userRepository Repository
}

func New(repository Repository) Validator {
	return Validator{
		userRepository: repository,
	}
}
