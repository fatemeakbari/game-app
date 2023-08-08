package main

import (
	adapter "gameapp/adapter/redisadapter"
	"gameapp/cfg"
	"gameapp/delivery/grpcserver/presence"
	presenceredis "gameapp/repository/redis/presence"
	presenceservice "gameapp/service/presence"
)

func main() {

	config := cfg.Load()

	adapter := adapter.New(config.Redis)

	presenceRep := presenceredis.New(adapter)
	service := presenceservice.New(presenceRep, config.Presence)
	server := presence.New(service)

	server.Start()
}
