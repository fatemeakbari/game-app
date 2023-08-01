package middleware

import (
	authservice "gameapp/service/auth"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthMiddleWare struct {
	authService authservice.Service
	authConfig  authservice.Config
}

func New(authService authservice.Service) AuthMiddleWare {
	return AuthMiddleWare{
		authService: authService,
	}
}

func (awm *AuthMiddleWare) AuthorizeToken() echo.MiddlewareFunc {

	return echojwt.WithConfig(echojwt.Config{
		ContextKey: "claims",
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return awm.authService.Parse(auth)
		},
	})
}
