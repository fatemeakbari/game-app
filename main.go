package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"messagingapp/cfg"
	"messagingapp/pkg/hashing"
	"messagingapp/repository/mysql"
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

	userService := userservice.Service{
		UserRepository: userRep,
		Hashing:        hashing.SHA256{},
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
		fmt.Sprintf(`{"authorized_token":"%s"`, response.AuthorizedToken),
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

	claims, err := parseJwtToken(token)

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

func parseJwtToken(tokenStr string) (userservice.Claims, error) {

	var claims userservice.Claims

	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(cfg.TokenSecretKey), nil
	})

	if err != nil {
		return userservice.Claims{}, errors.New("error in parsing token")
	}

	return claims, nil
}
