package cfg

import (
	"gameapp/repository/mysql"
	"gameapp/repository/redis"
	authservice "gameapp/service/auth"
)

type Config struct {
	DB    mysql.Config       `koanf:"DB"`
	Auth  authservice.Config `koanf:"auth"`
	Redis redis.Config       `koanf:"redis""`
}
