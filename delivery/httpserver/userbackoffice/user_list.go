package userbackofficehandler

import (
	entity "gameapp/entity/userbackoffice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) userListHandler(c echo.Context) error {

	response, err := h.userBackofficeService.UserList(entity.UserListRequest{})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()) //TODO
	}

	return c.JSON(http.StatusOK, response)
}
