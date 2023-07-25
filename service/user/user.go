package user

import (
	"errors"
	"fmt"
	"messagingapp/entity"
	"messagingapp/pkg/phonenumber"
	authservice "messagingapp/service/auth"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	FindUserByPhoneNumber(phoneNumber string) (entity.User, error)
	FindUserById(userId uint) (entity.User, error)
}

type Hashing interface {
	Hash(str string) string
}

type Service struct {
	UserRepository Repository
	Hashing        Hashing
	TokenGenerator authservice.JwtTokenGenerator
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
	Password    string
}

type RegisterResponse struct {
	entity.User
}

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type ProfileRequest struct {
	UserId uint
}

type ProfileResponse struct {
	Name string
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
			Password:    s.Hashing.Hash(req.Password),
		},
	)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("error is save user %w", err)
	}

	return RegisterResponse{user}, nil
}

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {

	user, err := s.UserRepository.FindUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return LoginResponse{}, err
	}

	if user.Password != s.Hashing.Hash(req.Password) {
		return LoginResponse{}, errors.New("phone number or password is wrong")
	}

	accToken, err := s.TokenGenerator.GenerateAccessToken(user)

	if err != nil {
		return LoginResponse{}, errors.New("unexpected error in generating access token")
	}

	refToken, err := s.TokenGenerator.GenerateRefreshToken(user)

	if err != nil {
		return LoginResponse{}, errors.New("unexpected error in generating refresh token")
	}

	return LoginResponse{AccessToken: accToken, RefreshToken: refToken}, nil

}

func (s *Service) Profile(req ProfileRequest) (ProfileResponse, error) {

	user, err := s.UserRepository.FindUserById(req.UserId)

	if err != nil {
		return ProfileResponse{}, err
	}

	return ProfileResponse{
		Name: user.Name,
	}, nil

}

func (s *Service) validPhoneNumber(phoneNumber string) (bool, error) {

	if !phonenumber.IsValid(phoneNumber) {
		return false, errors.New("phone number is not valid format")
	}

	if res, err := s.UserRepository.IsPhoneNumberUnique(phoneNumber); err != nil {
		return false, err
	} else if !res {
		return res, errors.New("phone number is duplicated")
	}

	return true, nil

}
