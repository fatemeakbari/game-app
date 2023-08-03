package userhandler

import (
	accesscontrolmodel "gameapp/model/accesscontrol"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Route(e *echo.Echo) {

	userGroup := e.Group("/users")
	userGroup.POST("/register", h.userRegisterHandler)
	userGroup.POST("/login", h.userLoginHandler)
	userGroup.GET("/profile", h.userProfileHandler, h.AuthorizeToken(), h.IsUserHasAccess(accesscontrolmodel.UserList))
}
