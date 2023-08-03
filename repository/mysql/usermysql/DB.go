package usermysql

import (
	"database/sql"
	"gameapp/repository/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DB struct {
	db *sql.DB
}

func New(cfg mysql.Config) *DB {

	db, err := sql.Open(cfg.Driver, cfg.BuildURL())
	if err != nil {
		log.Fatal("can not connect to sql driver", err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DB{db: db}
}
