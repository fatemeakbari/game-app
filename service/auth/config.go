package authservice

import "time"

type Config struct {
	TokenExpirationDuration time.Duration
	TokenRefreshDuration    time.Duration
	TokenSecretKey          string
}
