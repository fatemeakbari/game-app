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

func New() *DB {

	db, err := sql.Open("mysql", "root:12345@(localhost:3309)/messagingapp")
	if err != nil {
		log.Fatal("can not connect to sql driver", err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &DB{db: db}
}
