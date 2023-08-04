package userbackofficeentity

type UserListRequest struct {
}

type UserListResponse struct {
	UserInfos []UserInfo `json:"users"`
}
