package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "cake:secret@tcp(cake-store-db:3306)/cake_store?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
