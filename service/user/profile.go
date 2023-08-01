package user

import . "gameapp/entity/user"

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
