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
	middleware.ACLMiddleware
}

func New(userService userservice.Service,
	authService authservice.Service,
	authMW middleware.AuthMiddleWare,
	aclMW middleware.ACLMiddleware) *Handler {
	return &Handler{
		userService:    userService,
		authService:    authService,
		AuthMiddleWare: authMW,
		ACLMiddleware:  aclMW,
	}
}
