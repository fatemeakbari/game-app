package main

import (
	"context"
	"fmt"
	"gameapp/adapter/presenceadapter"
	adapter "gameapp/adapter/redisadapter"
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
	presenceredis "gameapp/repository/redis/presence"
	matchingservice "gameapp/service/matching"
	presenceservice "gameapp/service/presence"
	matchingvalidator "gameapp/validator/matching"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"

	accesscontrolservice "gameapp/service/accesscontrol"
	"gameapp/service/auth"

	userservice "gameapp/service/user"
	userbackofficeservice "gameapp/service/userbackoffice"

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

	adapter := adapter.New(config.Redis)
	redisDB := matchingredis.New(adapter)
	matchingValidator := matchingvalidator.New()
	presenceClient := presenceadapter.New("localhost:8086")
	matchingService := matchingservice.New(redisDB, matchingValidator, presenceClient)

	userBackOfficeService := userbackofficeservice.New(userRepository)

	authMW := middleware.New(authService)

	aclRepository := accesscontrolmysql.New(config.DB)
	aclService := accesscontrolservice.New(aclRepository)
	aclMW := middleware.NewACLMiddleware(aclService)

	presenceRep := presenceredis.New(adapter)
	presenceService := presenceservice.New(presenceRep, config.Presence)
	presenceMW := middleware.NewPresence(presenceService)

	userHandler := *userhandler.New(userService, authService, authMW, presenceMW)

	userBackOfficeHandler := *userbackofficehandler.New(userBackOfficeService, authMW, aclMW)

	matchingHandler := matchinghandler.New(matchingService, authMW, presenceMW)

	server := httpserver.Server{
		Router:                echo.New(),
		UserHandler:           userHandler,
		UserBackOfficeHandler: userBackOfficeHandler,
		MatchingHandler:       matchingHandler,
	}

	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	fmt.Println("received interrupt signal, shutting down gracefully..")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, config.HttpServer.ServerShutdownDuration)

	defer cancel()

	if err := server.Router.Shutdown(ctx); err != nil {
		fmt.Println("serve not shut down successfully")
	}
}
