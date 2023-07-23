package userservice

import (
	"errors"
	"fmt"
	"messagingapp/entity"
	"messagingapp/pkg/phonenumber"
)

type UserRepository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
}

type UserService struct {
	userRepo UserRepository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	entity.User
}

func (s *UserService) Register(req RegisterRequest) (RegisterResponse, error) {

	if res, err := s.validPhoneNumber(req.PhoneNumber); err != nil || !res {
		return RegisterResponse{}, fmt.Errorf("phone number must be valid %w", err)
	}

	if len(req.Name) == 0 {
		return RegisterResponse{}, errors.New("your name must not be empty")
	}

	user, err := s.userRepo.Register(
		entity.User{
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
		},
	)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("error is save user %w", err)
	}

	return RegisterResponse{user}, nil
}

func (s *UserService) validPhoneNumber(phoneNumber string) (bool, error) {

	if !phonenumber.IsValid(phoneNumber) {
		return false, nil
	}

	if res, err := s.userRepo.IsPhoneNumberUnique(phoneNumber); err != nil {
		return false, err
	} else {
		return res, nil
	}

}
