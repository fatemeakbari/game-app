package user

import (
	"gameapp/model"
	authservice "gameapp/service/auth"
	validator "gameapp/validator/user"
)

type Repository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	Register(user model.User) (model.User, error)
	FindUserByPhoneNumber(phoneNumber string) (model.User, error)
	FindUserById(userId uint) (model.User, error)
}

type Hashing interface {
	Hash(str string) string
}

type Service struct {
	UserRepository Repository
	Hashing        Hashing
	Validator      validator.Validator
	TokenGenerator authservice.JwtTokenGenerator
}
