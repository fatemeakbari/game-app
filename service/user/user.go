package user

import (
	"errors"
	. "gameapp/entity/userentity"
	"gameapp/model"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
	authservice "gameapp/service/auth"
	validator "gameapp/validator/user"
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
	Validator      validator.Validator
	TokenGenerator authservice.JwtTokenGenerator
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {

	if err := s.Validator.ValidateRegister(req); err != nil {
		return RegisterResponse{}, err
	}

	user, err := s.UserRepository.Register(
		model.User{
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Password:    s.Hashing.Hash(req.Password),
		},
	)

	if err != nil {
		return RegisterResponse{}, errorhandler.New().
			WithWrappedError(err).
			WithMessage(errormessage.InternalError).
			WithCodeStatus(errorcodestatus.InternalError).
			WithOperation("register")
	}

	return RegisterResponse{
		UserInfo: UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		}}, nil
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
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}, nil

}
