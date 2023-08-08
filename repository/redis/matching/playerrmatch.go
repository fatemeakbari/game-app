package matchingredis

import (
	"context"
	"fmt"
	matchingentity "gameapp/entity/matching"
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
	_, err := db.adapter.Client.ZAdd(context.Background(), key, goredis.Z{
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

func (db *DB) RemoveFromWaitingList(ctx context.Context, category model.Category, userIds ...uint) error {

	//TODO handle empty list
	if len(userIds) == 0 {
		return nil
	}
	key := fmt.Sprintf(waitingListNameFormat, category)

	//_, err := db.adapter.Client.ZRem(context.Background(), key, goredis.Z{
	//	Score:  float64(time.Now().UnixMilli()),
	//	Member: strconv.Itoa(int(userId)),
	//}).Result()

	members := make([]any, 0)

	for _, userID := range userIds {
		members = append(members, strconv.Itoa(int(userID)))
	}
	_, err := db.adapter.Client.ZRem(ctx, key, members).Result()

	if err != nil {
		return fmt.Errorf("cant not remove user to waiting list %w", err)
	}

	return nil
}

func (db *DB) GetWaitingPlayerByCategory(category model.Category) ([]matchingentity.WaitingMember, error) {

	key := getKey(category)

	zset, err := db.adapter.Client.ZRangeWithScores(context.Background(), key, 0, -1).Result()

	waitingList := make([]matchingentity.WaitingMember, 0)

	//TODO handle err
	if err != nil {
		return waitingList, err
	}
	for _, item := range zset {

		userId, _ := strconv.Atoi(item.Member.(string))

		waitingList = append(waitingList, matchingentity.WaitingMember{
			UserID:    uint(userId),
			Timestamp: int64(item.Score),
		})

	}

	return waitingList, nil

}

func getKey(category model.Category) string {
	return fmt.Sprintf(waitingListNameFormat, category)
}
