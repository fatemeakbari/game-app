package userhandler

import (
	entity "gameapp/entity/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) userLoginHandler(c echo.Context) error {

	var loginReq entity.LoginRequest

	if err := c.Bind(&loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request format is wrong")
	}

	response, err := h.userService.Login(loginReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)

}
