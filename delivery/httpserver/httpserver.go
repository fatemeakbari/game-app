package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	authservice "messagingapp/service/auth"
	"messagingapp/service/user"
	"net/http"
)

type Server struct {
	AuthService authservice.Service
	UserService user.Service
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

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

	//http.HandleFunc("/users/register", s.UserRegisterHandler)
	http.HandleFunc("/users/login", s.UserLoginHandler)
	http.HandleFunc("/users/profile", s.UserProfileHandler)
	http.ListenAndServe(":8080", nil)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
