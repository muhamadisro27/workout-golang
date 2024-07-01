package main

import (
	"golang-restful-api/helper"
	"golang-restful-api/injector"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server := injector.InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
