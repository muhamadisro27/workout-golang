package database

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
)


func TestGetConnection(t *testing.T) {
	db := GetConnection()

	defer db.Close()
}