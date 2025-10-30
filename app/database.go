package app

import (
	"database/sql"
	"golang-restfulapi/helper"

	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/belajar-golang-db?parseTime=true&loc=Asia%2FJakarta")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
