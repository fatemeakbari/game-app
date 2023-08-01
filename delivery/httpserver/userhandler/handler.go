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
}

func New(userService userservice.Service, authService authservice.Service, authMW middleware.AuthMiddleWare) *Handler {
	return &Handler{
		userService:    userService,
		authService:    authService,
		AuthMiddleWare: authMW,
	}
}
