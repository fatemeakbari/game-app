package main

import (
	"gameapp/cfg"
	"gameapp/delivery/httpserver"
	"gameapp/delivery/httpserver/middleware"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/pkg/hashing"
	"gameapp/repository/mysql/usermysql"
	"gameapp/service/auth"
	userservice "gameapp/service/user"
	uservalidator "gameapp/validator/user"
)

func main() {

	config := cfg.Load()

	userRepository := usermysql.New(config.DB)
	authService := authservice.New(config.Auth)
	userValidator := uservalidator.New(userRepository)

	userService := userservice.Service{
		UserRepository: userRepository,
		TokenGenerator: authService,
		Hashing:        hashing.SHA256{},
		Validator:      userValidator,
	}

	authMW := middleware.New(authService)

	userHandler := *userhandler.New(userService, authService, authMW)

	server := httpserver.Server{
		UserHandler: userHandler,
	}

	server.Serve()

}
