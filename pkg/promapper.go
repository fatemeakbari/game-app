package pkg

import (
	"gameapp/contract/goproto/presence"
	entity "gameapp/entity/presence"
)

func MapToProGetPresenceRequest(req entity.GetPresenceRequest) presence.GetPresenceRequest {

	userIds := make([]uint64, 0)

	for _, id := range req.UserIds {
		userIds = append(userIds, uint64(id))
	}

	return presence.GetPresenceRequest{
		UserIds: userIds,
	}
}
func MapToProGetPresenceResponse(res entity.GetPresenceResponse) *presence.GetPresenceResponse {

	infos := make([]*presence.PresenceInfo, 0)

	for _, info := range res.Infos {
		infos = append(infos, &presence.PresenceInfo{
			UserId:    uint64(info.UserId),
			Timestamp: uint64(info.Timestamp),
		})
	}

	return &presence.GetPresenceResponse{
		Infos: infos,
	}
}

func MapToGetPresenceRequest(req presence.GetPresenceRequest) *entity.GetPresenceRequest {
	userIs := make([]uint, 0)

	for _, id := range req.GetUserIds() {
		userIs = append(userIs, uint(id))
	}

	return &entity.GetPresenceRequest{
		UserIds: userIs,
	}
}
func MapToGetPresenceResponse(res *presence.GetPresenceResponse) *entity.GetPresenceResponse {

	infos := make([]entity.PresenceInfo, 0)

	for _, info := range res.Infos {

		infos = append(infos, entity.PresenceInfo{
			UserId:    uint(info.UserId),
			Timestamp: int64(info.Timestamp),
		})
	}

	return &entity.GetPresenceResponse{
		Infos: infos,
	}
}
