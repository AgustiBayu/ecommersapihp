package app

import (
	"EcommersAPIHP/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/ecommersapihp?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
