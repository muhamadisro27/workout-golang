package belajar_golang_db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	dsn := "root:123456@unix(/home/user/workout-golang/7.golang-mysql/mysql/data/mysql.sock)/belajar_golang_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")

	return db
}

func CreatePrepareStatement(ctx context.Context, db *sql.DB, scriptSQL string) (stmt *sql.Stmt) {
	stmt, err := db.PrepareContext(ctx, scriptSQL)
	if err != nil {
		panic(err)
	}

	return stmt
}
