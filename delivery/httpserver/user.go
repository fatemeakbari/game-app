package httpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	userservice "messagingapp/service/user"
	"net/http"
	"strings"
)

func (s *Server) UserRegisterHandler(c echo.Context) error {

	var registerReq userservice.RegisterRequest

	if err := c.Bind(&registerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request form is wrong")
	}

	response, err := s.UserService.Register(registerReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)
}

func (s *Server) UserLoginHandler(writer http.ResponseWriter, req *http.Request) {

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

	response, err := s.UserService.Login(loginReq)

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

func (s *Server) UserProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprintf(
			writer,
			"invalid method",
		)
		return
	}

	token := req.Header.Get("Authorization")

	token = strings.Replace(token, "Bearer ", "", 1)

	claims, err := s.AuthService.Parse(token)

	if err != nil {
		errors.New("token is wrong")
	}
	profileReq := userservice.ProfileRequest{
		UserId: claims.UserID,
	}

	response, err := s.UserService.Profile(profileReq)

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
