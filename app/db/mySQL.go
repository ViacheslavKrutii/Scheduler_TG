package db

import (
	"database/sql"
	"time"
)

type MysqlSchedulerState struct {
	db *sql.DB
}

func NewMysqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/menu")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func NewMysqlMenuState() *MysqlMenuState {
	return &MysqlMenuState{
		NewMysqlConnection(),
	}
}
