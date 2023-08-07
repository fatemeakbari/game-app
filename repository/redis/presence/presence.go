package presenceredis

import (
	"context"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
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
