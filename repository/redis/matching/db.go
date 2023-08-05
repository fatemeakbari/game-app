package matchingredis

import (
	"context"
	"fmt"
	"gameapp/repository/redis"
	goredis "github.com/redis/go-redis/v9"
)

type DB struct {
	db  *goredis.Client
	ctx context.Context
}

func New(cfg redis.Config) *DB {

	db := goredis.NewClient(&goredis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	return &DB{
		db:  db,
		ctx: context.Background(),
	}
}
