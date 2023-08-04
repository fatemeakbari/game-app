package userbackofficeservice

import "gameapp/model/usermodel"

type UserRepository interface {
	UserList() ([]usermodel.User, error)
}
type Service struct {
	userRepository UserRepository
}

func New(rep UserRepository) Service {

	return Service{
		userRepository: rep,
	}
}
