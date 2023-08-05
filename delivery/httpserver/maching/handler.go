package matchinghandler

import (
	"gameapp/delivery/httpserver/middleware"
	playermatchservice "gameapp/service/matching"
)

type Handler struct {
	service playermatchservice.Service
	middleware.AuthMiddleWare
}

func New(service playermatchservice.Service, auth middleware.AuthMiddleWare) Handler {

	return Handler{
		service:        service,
		AuthMiddleWare: auth,
	}
}
