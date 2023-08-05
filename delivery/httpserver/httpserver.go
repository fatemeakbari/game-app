package httpserver

import (
	"gameapp/delivery/httpserver/maching"
	backofficehandler "gameapp/delivery/httpserver/userbackoffice"
	"gameapp/delivery/httpserver/userhandler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	UserHandler           userhandler.Handler
	UserBackOfficeHandler backofficehandler.Handler
	PlayerMatchHandler    matchinghandler.Handler
}

func (s *Server) Serve() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	s.UserHandler.Route(e)
	s.UserBackOfficeHandler.Route(e)
	s.PlayerMatchHandler.Route(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
