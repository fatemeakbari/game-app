package middleware

import (
	"gameapp/entity/auth"
	entity "gameapp/entity/presence"
	presenceservice "gameapp/service/presence"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type PresenceMiddleWare struct {
	service presenceservice.Service
}

func NewPresence(service presenceservice.Service) PresenceMiddleWare {
	return PresenceMiddleWare{
		service: service,
	}
}

func (pmw *PresenceMiddleWare) UpsertPresence() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			claims := c.Get("claims").(auth.Claims)

			_, err := pmw.service.Upsert(c.Request().Context(),
				entity.UpsertPresenceRequest{
					UserId:    claims.UserID,
					Timestamp: time.Now().UnixMicro(),
				})

			//TODO error handing is not correct
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

			return next(c)
		}
	}
}
