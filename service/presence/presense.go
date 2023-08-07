package presenceservice

import (
	"context"
	"fmt"
	entity "gameapp/entity/presence"
	"time"
)

type Repository interface {
	Upsert(ctx context.Context, key string, timestamp int64, exp time.Duration) error
}

type Config struct {
	KeyPrefix              string        `koanf:"key_prefix"`
	PresenceExpireDuration time.Duration `koanf:"presence_expire_duration"`
}

type Service struct {
	rep Repository
	cfg Config
}

func New(rep Repository, cfg Config) Service {

	return Service{
		rep: rep,
		cfg: cfg,
	}
}
func (s *Service) Upsert(ctx context.Context, req entity.UpsertPresenceRequest) (entity.UpsertPresenceResponse, error) {

	key := fmt.Sprintf("%s:%d", s.cfg.KeyPrefix, req.UserId)

	err := s.rep.Upsert(ctx, key, req.Timestamp, s.cfg.PresenceExpireDuration)

	//if err != nil{
	//	return entity.UpsertPresenceResponse{}, err
	//}

	return entity.UpsertPresenceResponse{}, err
}
