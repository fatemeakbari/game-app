package cfg

import (
	"gameapp/repository/mysql"
	authservice "gameapp/service/auth"
)

type Config struct {
	DB   mysql.Config       `koanf:"DB"`
	Auth authservice.Config `koanf:"auth"`
}
