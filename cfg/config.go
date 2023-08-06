package cfg

import (
	"gameapp/repository/mysql"
	"gameapp/repository/redis"
	authservice "gameapp/service/auth"
	"time"
)

type HttpServer struct {
	Port                   int           `koanf:"port"`
	ServerShutdownDuration time.Duration `koanf:"server_shutdown_duration"`
}
type Config struct {
	HttpServer HttpServer         `koanf:"http_server"`
	DB         mysql.Config       `koanf:"DB"`
	Auth       authservice.Config `koanf:"auth"`
	Redis      redis.Config       `koanf:"redis""`
}
