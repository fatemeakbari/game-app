package matchingentity

import "gameapp/model"

type AddUserToWaitingListRequest struct {
	UserId   uint           `json:"-"`
	Category model.Category `json:"category"`
}

type AddUserToWaitingListResponse struct {
}
