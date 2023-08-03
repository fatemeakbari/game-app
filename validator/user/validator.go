package user

import (
	"gameapp/model/usermodel"
)

const (
	phoneNumberPattern = `^09\d{9}`
)

type Repository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	FindUserByPhoneNumber(phoneNumber string) (usermodel.User, error)
}

type Validator struct {
	userRepository Repository
}

func New(repository Repository) Validator {
	return Validator{
		userRepository: repository,
	}
}
