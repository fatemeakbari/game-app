package matchingservice

import (
	"context"
	"fmt"
	entity "gameapp/entity/matching"
	presenceentity "gameapp/entity/presence"
	"gameapp/model"
	"time"
)

type Repository interface {
	AddUserToWaitingList(userId uint, category model.Category) error
	RemoveFromWaitingList(ctx context.Context, category model.Category, userIds ...uint) error
	GetWaitingPlayerByCategory(category model.Category) ([]entity.WaitingMember, error)
}

type Validator interface {
	ValidateAddUserToWaitingListRequest(request entity.AddUserToWaitingListRequest) error
}

type PresenceClient interface {
	GetPresence(ctx context.Context, request presenceentity.GetPresenceRequest) (*presenceentity.GetPresenceResponse, error)
}

type Service struct {
	rep            Repository
	validator      Validator
	presenceClient PresenceClient
}

func New(rep Repository, validator Validator, client PresenceClient) Service {

	return Service{
		rep:            rep,
		validator:      validator,
		presenceClient: client,
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

func (s *Service) MatchWaitingPlayer(ctx context.Context, category model.Category) error {

	var waitingList, err = s.rep.GetWaitingPlayerByCategory(category)

	if err != nil {
		fmt.Println("waitingList err ", err)
		return err
	}

	fmt.Println("waitingList", waitingList)

	userIds := make([]uint, 0)

	for _, item := range waitingList {
		userIds = append(userIds, item.UserID)
	}

	req, err := s.presenceClient.GetPresence(ctx, presenceentity.GetPresenceRequest{
		UserIds: userIds,
	})

	if err != nil {
		return err
	}

	finalUserIDs := getValidUserIDToAddToTheGame(req.Infos, userIds)

	finalUserIDsLen := len(finalUserIDs)

	if finalUserIDsLen%2 == 1 {
		finalUserIDs = finalUserIDs[:finalUserIDsLen-1]
	}

	//TODO can run in goroutine
	func() {
		for i := 0; i < len(finalUserIDs); i += 2 {
			//TODO send matching player event
			fmt.Printf("matching user %d and user %d\n", finalUserIDs[i], finalUserIDs[i+1])
		}
	}()

	//TODO can run in goroutine
	err = s.rep.RemoveFromWaitingList(ctx, category, finalUserIDs...)

	if err != nil {
		fmt.Println("")
		return err
	}
	return nil
}

func getValidUserIDToAddToTheGame(presenceInfos []presenceentity.PresenceInfo, waitingUserIds []uint) []uint {

	finalUserIDs := make([]uint, 0)

	presenceInfosMap := make(map[uint]int64, 0)

	for _, info := range presenceInfos {

		presenceInfosMap[info.UserId] = info.Timestamp
	}

	for _, userID := range waitingUserIds {

		timestamp, ok := presenceInfosMap[userID]

		if ok {
			//TODO expire time must be config
			if timestamp >= time.Now().Add(-60*time.Minute).UnixMicro() {
				finalUserIDs = append(finalUserIDs, userID)
			}
		}
	}
	return finalUserIDs
}
