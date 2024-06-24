package belajar_golang_db

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	fmt.Println("Test lagi")
}

func TestOpenConnection(t *testing.T) {
	db := GetConnection()

	defer db.Close()
}
