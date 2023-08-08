package cfg

import (
	redis "gameapp/adapter/redisadapter"
	"gameapp/repository/mysql"
	auth "gameapp/service/auth"
	presence "gameapp/service/presence"
	"time"
)

type HttpServer struct {
	Port                   int           `koanf:"port"`
	ServerShutdownDuration time.Duration `koanf:"server_shutdown_duration"`
}
type Config struct {
	HttpServer HttpServer      `koanf:"http_server"`
	DB         mysql.Config    `koanf:"DB"`
	Auth       auth.Config     `koanf:"auth"`
	Redis      redis.Config    `koanf:"redis""`
	Presence   presence.Config `koanf:"presence"`
}
