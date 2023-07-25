package httpserver

import (
	authservice "messagingapp/service/auth"
	"messagingapp/service/user"
	"net/http"
)

type Server struct {
	AuthService authservice.Service
	UserService user.Service
}

func (s *Server) Serve() {

	http.HandleFunc("/users/register", s.UserRegisterHandler)
	http.HandleFunc("/users/login", s.UserLoginHandler)
	http.HandleFunc("/users/profile", s.UserProfileHandler)
	http.ListenAndServe(":8080", nil)
}
