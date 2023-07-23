package user

import (
	"errors"
	"fmt"
	"messagingapp/entity"
	"messagingapp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
}

type Service struct {
	UserRepository Repository
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
}

type RegisterResponse struct {
	entity.User
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {

	if res, err := s.validPhoneNumber(req.PhoneNumber); err != nil || !res {
		return RegisterResponse{}, err
	}

	if len(req.Name) == 0 {
		return RegisterResponse{}, errors.New("your name must not be empty")
	}

	user, err := s.UserRepository.Register(
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

func (s *Service) validPhoneNumber(phoneNumber string) (bool, error) {

	if !phonenumber.IsValid(phoneNumber) {
		return false, errors.New("phone number is not valid format")
	}

	if res, err := s.UserRepository.IsPhoneNumberUnique(phoneNumber); err != nil {
		return false, err
	} else {
		return res, nil
	}

}
