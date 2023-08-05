package matchinghandler

import (
	"gameapp/delivery/httpserver/middleware"
	matchingservice "gameapp/service/matching"
)

type Handler struct {
	service matchingservice.Service
	middleware.AuthMiddleWare
}

func New(service matchingservice.Service, auth middleware.AuthMiddleWare) Handler {

	return Handler{
		service:        service,
		AuthMiddleWare: auth,
	}
}
