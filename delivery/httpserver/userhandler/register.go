package userhandler

import (
	entity "gameapp/entity/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) userRegisterHandler(c echo.Context) error {

	var registerReq entity.RegisterRequest

	if err := c.Bind(&registerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request form is wrong")
	}

	response, err := h.userService.Register(registerReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusCreated, response)
}
