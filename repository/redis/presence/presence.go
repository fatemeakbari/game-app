package presenceredis

import (
	"context"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"strconv"
	"time"
)

func (db *DB) Upsert(ctx context.Context, key string, timestamp int64, exp time.Duration) error {

	_, err := db.adapter.Client.Set(ctx, key, timestamp, exp).Result()

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithCodeStatus(errorcodestatus.InternalError).
			WithMessage(err.Error())
	}

	return nil
}
func (db *DB) GetPresence(ctx context.Context, keys []string) (map[string]uint, error) {

	result := make(map[string]uint)

	//TODO how handle empty key list?
	if len(keys) == 0 {
		return result, nil
	}

	res, err := db.adapter.Client.MGet(ctx, keys...).Result()

	//TODO handel err
	if err != nil {
		return nil, err
	}

	for idx, val := range res {

		if val != nil {
			val_, _ := strconv.Atoi(val.(string))
			result[keys[idx]] = uint(val_)
		}
	}

	return result, nil
}
