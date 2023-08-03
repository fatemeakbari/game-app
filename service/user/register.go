package user

import (
	. "gameapp/entity/user"
	"gameapp/model/usermodel"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
)

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {

	if err := s.Validator.ValidateRegisterRequest(req); err != nil {
		return RegisterResponse{}, err
	}

	user, err := s.UserRepository.Register(
		usermodel.User{
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
