package golangorm

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dialect := "root:123456@tcp(127.0.0.1:3306)/golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dialect), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	PanicIfError(err)

	sqlDB, err := db.DB()
	PanicIfError(err)

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	defer fmt.Println("Successfully connect to database !")

	return db
}

func CloseConnection(db *gorm.DB) {

	dbInstance, err := db.DB()
	PanicIfError(err)

	err = dbInstance.Close()
	PanicIfError(err)

	fmt.Println("Connection Closed")
}
