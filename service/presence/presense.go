package presenceservice

import (
	"context"
	"fmt"
	entity "gameapp/entity/presence"
	"time"
)

type Repository interface {
	Upsert(ctx context.Context, key string, timestamp int64, exp time.Duration) error
	GetPresence(ctx context.Context, userIds []string) (map[string]uint, error)
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

	err := s.rep.Upsert(ctx, s.getKey(req.UserId), req.Timestamp, s.cfg.PresenceExpireDuration)

	//if err != nil{
	//	return entity.UpsertPresenceResponse{}, err
	//}

	return entity.UpsertPresenceResponse{}, err
}

func (s *Service) GetPresence(ctx context.Context, req entity.GetPresenceRequest) (entity.GetPresenceResponse, error) {

	keys := make([]string, 0)
	userId2Key := make(map[string]uint)

	for _, userId := range req.UserIds {
		key := s.getKey(userId)
		userId2Key[key] = userId
		keys = append(keys, key)
	}

	result, err := s.rep.GetPresence(ctx, keys)

	//TODO handle err
	if err != nil {
		return entity.GetPresenceResponse{}, err
	}

	infos := make([]entity.PresenceInfo, 0)
	for k, v := range result {

		infos = append(infos, entity.PresenceInfo{
			UserId:    userId2Key[k],
			Timestamp: int64(v),
		})
	}

	return entity.GetPresenceResponse{
		Infos: infos,
	}, err
}

func (s *Service) getKey(userId uint) string {

	return fmt.Sprintf("%s:%d", s.cfg.KeyPrefix, userId)
}
