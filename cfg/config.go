package cfg

import (
	"gameapp/repository/mysql"
	authservice "gameapp/service/auth"
	"time"
)

const (
	TokenExpirationDuration = time.Hour * 24 * 7
	TokenRefreshDuration    = time.Hour * 24 * 30
	TokenSecretKey          = "mdnfkfsdfkhsdfjaslsfdsfsf"
)

type Config struct {
	DBConfig   mysql.Config
	AuthConfig authservice.Config
}
