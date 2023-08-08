package main

import (
	"fmt"
	"gameapp/adapter/presenceadapter"
	adapter "gameapp/adapter/redisadapter"
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
	adapter := adapter.New(config.Redis)
	redisDB := matchingredis.New(adapter)
	matchingValidator := matchingvalidator.New()
	presenceClient := presenceadapter.New("localhost:8086")
	matchingService := matchingservice.New(redisDB, matchingValidator, presenceClient)

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
