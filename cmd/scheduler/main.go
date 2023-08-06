package main

import (
	"fmt"
	"gameapp/cfg"
	matchingredis "gameapp/repository/redis/matching"
	"gameapp/scheduler"
	matchingservice "gameapp/service/matching"
	matchingvalidator "gameapp/validator/matching"
	"os"
	"os/signal"
)

func main() {

	config := cfg.Load()
	redisDB := matchingredis.New(config.Redis)
	matchingValidator := matchingvalidator.New()
	matchingService := matchingservice.New(redisDB, matchingValidator)

	sch := scheduler.New(matchingService)

	done := make(chan bool)

	go sch.Start(done)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Println("received interrupt signal, shutting down gracefully..")

	fmt.Println("stopping scheduler")
	//done <- true

}
