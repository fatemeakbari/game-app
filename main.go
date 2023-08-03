package main

import (
	"gameapp/cfg"
	"gameapp/delivery/httpserver"
	"gameapp/delivery/httpserver/middleware"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/pkg/hashing"
	accesscontrolmysql "gameapp/repository/mysql/accesscontrol"
	"gameapp/repository/mysql/usermysql"
	accesscontrolservice "gameapp/service/accesscontrol"
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

	aclRepository := accesscontrolmysql.New(config.DB)
	aclService := accesscontrolservice.New(aclRepository)
	aclMW := middleware.NewACLMiddleware(aclService)

	userHandler := *userhandler.New(userService, authService, authMW, aclMW)

	server := httpserver.Server{
		UserHandler: userHandler,
	}

	server.Serve()

}
