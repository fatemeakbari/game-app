package userbackofficeservice

import (
	entity "gameapp/entity/userbackoffice"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
)

func (s *Service) UserList(request entity.UserListRequest) (entity.UserListResponse, error) {

	users, err := s.userRepository.UserList()

	if err != nil {
		return entity.UserListResponse{}, errorhandler.New().
			WithWrappedError(err).
			WithCodeStatus(errorcodestatus.InternalError).
			WithMessage(errormessage.InternalError)
	}

	userInfos := make([]entity.UserInfo, 0)

	for _, user := range users {

		userInfos = append(userInfos, entity.UserInfo{
			ID:          user.ID,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		})
	}
	return entity.UserListResponse{UserInfos: userInfos}, nil
}
