package authservice

import "time"

type Config struct {
	TokenAccessDuration  time.Duration `koanf:"token_access_duration"`
	TokenRefreshDuration time.Duration `koanf:"token_refresh_duration"`
	TokenSecretKey       string        `koanf:"token_secret_key"`
}
