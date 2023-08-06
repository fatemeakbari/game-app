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
	Router                *echo.Echo
	UserHandler           userhandler.Handler
	UserBackOfficeHandler backofficehandler.Handler
	MatchingHandler       matchinghandler.Handler
}

func (s *Server) Serve() {

	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.UserHandler.Route(s.Router)
	s.UserBackOfficeHandler.Route(s.Router)
	s.MatchingHandler.Route(s.Router)

	// Start server
	//TODO server port must be confined
	s.Router.Logger.Fatal(s.Router.Start(":8080"))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
