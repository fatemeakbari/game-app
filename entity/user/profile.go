package user

type ProfileRequest struct {
	UserId uint `json:"user_id"`
}

type ProfileResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
