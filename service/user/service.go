package user

import (
	"gameapp/model/usermodel"
	authservice "gameapp/service/auth"
	validator "gameapp/validator/user"
)

type Repository interface {
	IsPhoneNumberExist(phoneNumber string) (bool, error)
	Register(user usermodel.User) (usermodel.User, error)
	FindUserByPhoneNumber(phoneNumber string) (usermodel.User, error)
	FindUserById(userId uint) (usermodel.User, error)
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
