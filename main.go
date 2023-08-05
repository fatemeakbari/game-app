package main

import (
	"gameapp/cfg"
	"gameapp/delivery/httpserver"
	matchinghandler "gameapp/delivery/httpserver/maching"
	"gameapp/delivery/httpserver/middleware"
	userbackofficehandler "gameapp/delivery/httpserver/userbackoffice"
	"gameapp/delivery/httpserver/userhandler"
	"gameapp/pkg/hashing"
	accesscontrolmysql "gameapp/repository/mysql/accesscontrol"
	"gameapp/repository/mysql/usermysql"
	matchingredis "gameapp/repository/redis/matching"
	accesscontrolservice "gameapp/service/accesscontrol"
	"gameapp/service/auth"
	matchingservice "gameapp/service/matching"
	userservice "gameapp/service/user"
	userbackofficeservice "gameapp/service/userbackoffice"
	matchingvalidator "gameapp/validator/matching"
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

	redisDB := matchingredis.New(config.Redis)
	matchingValidator := matchingvalidator.New()
	platerMatchService := matchingservice.New(redisDB, matchingValidator)

	userBackOfficeService := userbackofficeservice.New(userRepository)

	authMW := middleware.New(authService)

	aclRepository := accesscontrolmysql.New(config.DB)
	aclService := accesscontrolservice.New(aclRepository)
	aclMW := middleware.NewACLMiddleware(aclService)

	userHandler := *userhandler.New(userService, authService, authMW)

	userBackOfficeHandler := *userbackofficehandler.New(userBackOfficeService, authMW, aclMW)

	matchingHandler := matchinghandler.New(platerMatchService, authMW)

	server := httpserver.Server{
		UserHandler:           userHandler,
		UserBackOfficeHandler: userBackOfficeHandler,
		MatchingHandler:       matchingHandler,
	}

	server.Serve()

}
