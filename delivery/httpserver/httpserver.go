package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authservice "messagingapp/service/auth"
	"messagingapp/service/userservice"
	"net/http"
)

type Server struct {
	AuthService authservice.Service
	UserService userservice.Service
}

func (s *Server) Serve() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	userGroup := e.Group("/users")
	userGroup.POST("/register", s.UserRegisterHandler)
	userGroup.POST("/login", s.UserLoginHandler)
	userGroup.GET("/profile", s.UserProfileHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
