package matchingredis

import (
	"fmt"
	"gameapp/model"
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
		return fmt.Errorf("cant not add user to waiting list %w", err)
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
