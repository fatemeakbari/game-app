package userentity

type LoginRequest struct {
	PhoneNumber string
	Password    string
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
