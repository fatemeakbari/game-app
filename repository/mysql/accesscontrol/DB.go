package accesscontrolmysql

import (
	"database/sql"
	"gameapp/repository/mysql"
)

type DB struct {
	db *sql.DB
}

func New(cfg mysql.Config) *DB {
	conn, err := sql.Open(cfg.Driver, cfg.BuildURL())

	if err != nil {
		panic(err)
	}

	return &DB{db: conn}
}
