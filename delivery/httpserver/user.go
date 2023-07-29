package httpserver

import (
	"github.com/labstack/echo/v4"
	entity "messagingapp/entity/userentity"
	"net/http"
	"strings"
)

func (s *Server) UserRegisterHandler(c echo.Context) error {

	var registerReq entity.RegisterRequest

	if err := c.Bind(&registerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request form is wrong")
	}

	response, err := s.UserService.Register(registerReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)
}

func (s *Server) UserLoginHandler(c echo.Context) error {

	var loginReq entity.LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request format is wrong")
	}

	response, err := s.UserService.Login(loginReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)

}

func (s *Server) UserProfileHandler(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	token = strings.Replace(token, "Bearer ", "", 1)

	claims, err := s.AuthService.Parse(token)

	profileReq := entity.ProfileRequest{
		UserId: claims.UserID,
	}

	response, err := s.UserService.Profile(profileReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)

}
