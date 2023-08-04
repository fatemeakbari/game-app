package userbackofficehandler

import (
	accesscontrolmodel "gameapp/model/accesscontrol"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Route(e *echo.Echo) {

	group := e.Group("/backoffice/users", h.AuthorizeToken(), h.IsUserHasAccess(accesscontrolmodel.UserList))

	group.GET("/", h.userListHandler)
}
