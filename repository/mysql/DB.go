package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DB struct {
	db *sql.DB
}

func New(cfg Config) *DB {

	db, err := sql.Open(cfg.Driver, cfg.buildURL())
	if err != nil {
		log.Fatal("can not connect to sql driver", err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DB{db: db}
}
