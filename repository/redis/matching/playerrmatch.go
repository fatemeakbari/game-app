package matchingredis

import (
	"fmt"
	"gameapp/model"
	"gameapp/pkg/errorhandler"
	"gameapp/pkg/errorhandler/errorcodestatus"
	"gameapp/pkg/errorhandler/errormessage"
	goredis "github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const waitingListNameFormat = `waitingList:%s`

func (db *DB) AddUserToWaitingList(userId uint, category model.Category) error {

	key := fmt.Sprintf(waitingListNameFormat, category)
	_, err := db.db.ZAdd(db.ctx, key, goredis.Z{
		Score:  float64(time.Now().UnixMilli()),
		Member: strconv.Itoa(int(userId)),
	}).Result()

	if err != nil {
		return errorhandler.New().
			WithWrappedError(err).
			WithOperation("repository.AddUserToWaitingList").
			WithMessage(errormessage.InternalError).
			WithCodeStatus(errorcodestatus.InternalError)
	}

	return nil

}

func (db *DB) RemoveUserFromWaitingList(userId uint, category model.Category) error {

	key := fmt.Sprintf(waitingListNameFormat, category)

	_, err := db.db.ZRem(db.ctx, key, goredis.Z{
		Score:  float64(time.Now().UnixMilli()),
		Member: strconv.Itoa(int(userId)),
	}).Result()

	if err != nil {
		return fmt.Errorf("cant not remove user to waiting list %w", err)
	}

	return nil
}

func (db *DB) GetWaitingPlayerByCategory(category model.Category) {

	key := getKey(category)

	zset, err := db.db.ZRangeWithScores(db.ctx, key, 0, -1).Result()

	if err != nil {
		panic(err)
	}
	if len(zset)%2 == 1 {
		zset = zset[:len(zset)-1]
	}

	for i := 0; i < len(zset)-1; i += 2 {

		item1 := zset[i]
		item2 := zset[i+1]
		fmt.Printf("matching user %s and user %s\n", item1.Member, item2.Member)
		db.db.ZRemRangeByScore(db.ctx, key, strconv.Itoa(int(item1.Score)), strconv.Itoa(int(item2.Score)))
	}

}

func getKey(category model.Category) string {
	return fmt.Sprintf(waitingListNameFormat, category)
}
