package presenceredis

import (
	adapter "gameapp/adapter/redis"
)

type DB struct {
	adapter adapter.Adapter
}

func New(adapter *adapter.Adapter) *DB {

	return &DB{
		adapter: *adapter,
	}
}
