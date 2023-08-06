package matchingservice

import (
	entity "gameapp/entity/matching"
	"gameapp/model"
)

type Repository interface {
	AddUserToWaitingList(userId uint, category model.Category) error
	RemoveUserFromWaitingList(userId uint, category model.Category) error
	GetWaitingPlayerByCategory(category model.Category)
}

type Validator interface {
	ValidateAddUserToWaitingListRequest(request entity.AddUserToWaitingListRequest) error
}

type Service struct {
	rep       Repository
	validator Validator
}

func New(rep Repository, validator Validator) Service {

	return Service{
		rep:       rep,
		validator: validator,
	}
}

func (s *Service) AddUserToWaitingList(req entity.AddUserToWaitingListRequest) (entity.AddUserToWaitingListResponse, error) {

	err := s.validator.ValidateAddUserToWaitingListRequest(req)

	if err != nil {
		return entity.AddUserToWaitingListResponse{}, err
	}

	err = s.rep.AddUserToWaitingList(req.UserId, req.Category)

	if err != nil {
		return entity.AddUserToWaitingListResponse{}, err
	}

	return entity.AddUserToWaitingListResponse{}, nil
}

func (s *Service) MatchWaitingPlayer(category model.Category) {

	s.rep.GetWaitingPlayerByCategory(category)
}
