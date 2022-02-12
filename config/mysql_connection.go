package config

import (
	"database/sql"
	"time"

	"github.com/andiahmads/go-restfulAPI/helper"
)

func SetupConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:endi@tcp(localhost:3306)/go_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}
