package matchingredis

import (
	adapter "gameapp/adapter/redisadapter"
)

type DB struct {
	adapter adapter.Adapter
}

func New(adapter *adapter.Adapter) *DB {

	return &DB{
		adapter: *adapter,
	}
}
