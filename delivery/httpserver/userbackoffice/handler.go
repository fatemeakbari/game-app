package userbackofficehandler

import (
	"gameapp/delivery/httpserver/middleware"
	userbackofficeservice "gameapp/service/userbackoffice"
)

type Handler struct {
	userBackofficeService userbackofficeservice.Service
	middleware.AuthMiddleWare
	middleware.ACLMiddleware
}

func New(service userbackofficeservice.Service, auth middleware.AuthMiddleWare, acl middleware.ACLMiddleware) *Handler {

	return &Handler{
		userBackofficeService: service,
		AuthMiddleWare:        auth,
		ACLMiddleware:         acl,
	}
}
