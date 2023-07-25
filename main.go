package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"messagingapp/cfg"
	"messagingapp/pkg/hashing"
	"messagingapp/repository/mysql"
	authservice "messagingapp/service/auth"
	userservice "messagingapp/service/user"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)
	http.HandleFunc("/users/profile", userProfileHandler)
	http.ListenAndServe(":8080", nil)
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		fmt.Fprintf(
			writer,
			"invalid method",
		)
		return
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in reading request %w"}`, err),
		)

		return
	}

	var ureq userservice.RegisterRequest
	if err := json.Unmarshal(body, &ureq); err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in parsing request %w"}`, err),
		)

		return
	}

	userRep := mysql.New()

	userService := userservice.Service{
		UserRepository: userRep,
		Hashing:        hashing.SHA256{},
	}

	response, err := userService.Register(ureq)

	if err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in register", details":"%s"}`, err),
		)
		return
	}

	byteRes, _ := json.Marshal(response)
	fmt.Fprintf(
		writer,
		string(byteRes),
	)

}

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		fmt.Fprintf(
			writer,
			"invalid method",
		)
		return
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in reading request %w"}`, err),
		)

		return
	}

	var loginReq userservice.LoginRequest

	if err := json.Unmarshal(body, &loginReq); err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in parsing request %w"}`, err),
		)

		return
	}

	userRep := mysql.New()
	generateTokenService := authservice.New(cfg.TokenSecretKey, cfg.TokenExpirationDuration, cfg.TokenRefreshDuration)

	userService := userservice.Service{
		UserRepository: userRep,
		Hashing:        hashing.SHA256{},
		TokenGenerator: generateTokenService,
	}

	response, err := userService.Login(loginReq)

	if err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in register", details":"%s"}`, err),
		)
		return
	}

	fmt.Fprintf(
		writer,
		fmt.Sprintf(`{"access_token":"%s", "refresh_token":"%s"}`, response.AccessToken, response.RefreshToken),
	)

}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(
			writer,
			"invalid method",
		)
		return
	}

	token := req.Header.Get("Authorization")

	token = strings.Replace(token, "Bearer ", "", 1)

	parseTokenService := authservice.New(cfg.TokenSecretKey, cfg.TokenExpirationDuration, cfg.TokenRefreshDuration)

	claims, err := parseTokenService.Parse(token)

	if err != nil {
		errors.New("token is wrong")
	}
	profileReq := userservice.ProfileRequest{
		UserId: claims.UserID,
	}

	userRep := mysql.New()

	userService := userservice.Service{
		UserRepository: userRep,
		Hashing:        hashing.SHA256{},
	}

	response, err := userService.Profile(profileReq)

	if err != nil {
		fmt.Fprintf(
			writer,
			fmt.Sprintf(`{"error": "error in register", details":"%s"}`, err),
		)
		return
	}

	strRes, _ := json.Marshal(response)

	fmt.Fprintf(writer, string(strRes))
}
