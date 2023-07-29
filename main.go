package main

import (
	"messagingapp/cfg"
	"messagingapp/delivery/httpserver"
	"messagingapp/pkg/hashing"
	"messagingapp/repository/mysql"
	"messagingapp/service/auth"
	userservice "messagingapp/service/userservice"
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

	userService := userservice.Service{
		UserRepository: userRepository,
		TokenGenerator: authService,
		Hashing:        hashing.SHA256{},
	}

	server := httpserver.Server{
		UserService: userService,
		AuthService: authService,
	}

	server.Serve()

}
