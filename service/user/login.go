package user

import (
	"errors"
	"fmt"
	. "gameapp/entity/user"
)

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {

	err := s.Validator.ValidateLoginRequest(req)

	if err != nil {
		return LoginResponse{}, err
	}

	user, _ := s.UserRepository.FindUserByPhoneNumber(req.PhoneNumber)

	if user.Password != s.Hashing.Hash(req.Password) {
		return LoginResponse{}, errors.New("phone number or password is wrong")
	}

	accToken, err := s.TokenGenerator.GenerateAccessToken(user)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	refToken, err := s.TokenGenerator.GenerateRefreshToken(user)

	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return LoginResponse{AccessToken: accToken, RefreshToken: refToken}, nil

}
