package database

import (
	"database/sql"
	"fmt"
	"golang-restful-api/helper"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetConnection() *sql.DB {
	// err := godotenv.Load("../../.env")
	// helper.PanicIfError(err)

	dsn := "root:123456@unix(/home/user/workout-golang/12.golang-restfulapi/mysql/data/mysql.sock)/belajar_golang_restful_api?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	helper.PanicIfError(err)

	fmt.Println("Successfully connected to the database!")

	return db
}
