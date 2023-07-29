package userentity

type ProfileRequest struct {
	UserId uint
}

type ProfileResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
