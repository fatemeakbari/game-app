package presenceentity

type GetPresenceRequest struct {
	UserIds []uint
}

type GetPresenceResponse struct {
	Infos []PresenceInfo
}
