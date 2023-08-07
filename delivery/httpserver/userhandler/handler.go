package userhandler

import (
	"gameapp/delivery/httpserver/middleware"
	authservice "gameapp/service/auth"
	userservice "gameapp/service/user"
)

type Handler struct {
	userService userservice.Service
	authService authservice.Service
	middleware.AuthMiddleWare
	middleware.PresenceMiddleWare
}

func New(userService userservice.Service,
	authService authservice.Service,
	authMW middleware.AuthMiddleWare, presence middleware.PresenceMiddleWare) *Handler {
	return &Handler{
		userService:        userService,
		authService:        authService,
		AuthMiddleWare:     authMW,
		PresenceMiddleWare: presence,
	}
}
