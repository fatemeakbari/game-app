package userservice

import (
	"errors"
	"fmt"
	entity "messagingapp/entity/userentity"
	"messagingapp/model"
	"messagingapp/pkg/phonenumber"
	authservice "messagingapp/service/auth"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
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
	TokenGenerator authservice.JwtTokenGenerator
}

func (s *Service) Register(req entity.RegisterRequest) (entity.RegisterResponse, error) {

	if res, err := s.validPhoneNumber(req.PhoneNumber); err != nil || !res {
		return entity.RegisterResponse{}, err
	}

	if len(req.Name) == 0 {
		return entity.RegisterResponse{}, errors.New("your name must not be empty")
	}

	user, err := s.UserRepository.Register(
		model.User{
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Password:    s.Hashing.Hash(req.Password),
		},
	)

	if err != nil {
		return entity.RegisterResponse{}, fmt.Errorf("error is save user %w", err)
	}

	return entity.RegisterResponse{
		UserInfo: entity.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		}}, nil
}

func (s *Service) Login(req entity.LoginRequest) (entity.LoginResponse, error) {

	user, err := s.UserRepository.FindUserByPhoneNumber(req.PhoneNumber)

	if err != nil {
		return entity.LoginResponse{}, err
	}

	if user.Password != s.Hashing.Hash(req.Password) {
		return entity.LoginResponse{}, errors.New("phone number or password is wrong")
	}

	accToken, err := s.TokenGenerator.GenerateAccessToken(user)

	if err != nil {
		return entity.LoginResponse{}, errors.New("unexpected error in generating access token")
	}

	refToken, err := s.TokenGenerator.GenerateRefreshToken(user)

	if err != nil {
		return entity.LoginResponse{}, errors.New("unexpected error in generating refresh token")
	}

	return entity.LoginResponse{AccessToken: accToken, RefreshToken: refToken}, nil

}

func (s *Service) Profile(req entity.ProfileRequest) (entity.ProfileResponse, error) {

	user, err := s.UserRepository.FindUserById(req.UserId)

	if err != nil {
		return entity.ProfileResponse{}, err
	}

	return entity.ProfileResponse{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
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
