package main

import (
	"encoding/json"
	"fmt"
	"io"
	"messagingapp/pkg/hashing"
	"messagingapp/repository/mysql"
	userservice "messagingapp/service/user"
	"net/http"
)

func main() {

	http.HandleFunc("/users/register", userRegisterHandler)
	http.HandleFunc("/users/login", userLoginHandler)
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
