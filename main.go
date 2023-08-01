package main

import (
	"gameapp/cfg"
	"gameapp/delivery/httpserver"
	"gameapp/delivery/httpserver/middleware"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/pkg/hashing"
	"gameapp/repository/mysql"
	"gameapp/service/auth"
	userservice "gameapp/service/user"
	uservalidator "gameapp/validator/user"
	"time"
)

func main() {

	config := cfg.Config{
		DBConfig: mysql.Config{
			Driver: "mysql",
			Name:   "messagingapp",
			User:   "root",
			Pass:   "12345",
			Host:   "localhost",
			Port:   "3309",
		},

		AuthConfig: authservice.Config{
			TokenExpirationDuration: time.Hour * 24 * 7,
			TokenRefreshDuration:    time.Hour * 24 * 30,
			TokenSecretKey:          "mdnfkfsdfkhsdfjaslsfdsfsf",
		},
	}
	userRepository := mysql.New(config.DBConfig)
	authService := authservice.New(config.AuthConfig)
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
