package accesscontrolservice

import (
	model "gameapp/model/accesscontrol"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
)

type ACLRepository interface {
	GetUserPermissionTitles(userId, roleId uint) ([]model.PermissionTitle, error)
}

type Service struct {
	repository ACLRepository
}

func New(rep ACLRepository) Service {
	return Service{
		repository: rep,
	}
}

func (s *Service) IsUserHasAccess(userId, role uint, titles []model.PermissionTitle) (bool, error) {

	op := "IsUserHasAccess"
	permissionTitles, err := s.repository.GetUserPermissionTitles(userId, role)

	if err != nil {
		return false, errorhandler.New().
			WithWrappedError(err).
			WithOperation(op).
			WithMessage(errormessage.ForbiddenMessage).
			WithCodeStatus(errorcodestatus.Forbidden)
	}

	for _, tout := range titles {

		for _, tin := range permissionTitles {
			if tin == tout {
				return true, nil
			}
		}
	}

	return false, nil
}
