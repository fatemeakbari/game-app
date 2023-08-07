package matchinghandler

import (
	"gameapp/delivery/httpserver/middleware"
	matchingservice "gameapp/service/matching"
)

type Handler struct {
	service matchingservice.Service
	middleware.AuthMiddleWare
	middleware.PresenceMiddleWare
}

func New(service matchingservice.Service, auth middleware.AuthMiddleWare, presence middleware.PresenceMiddleWare) Handler {

	return Handler{
		service:            service,
		AuthMiddleWare:     auth,
		PresenceMiddleWare: presence,
	}
}
