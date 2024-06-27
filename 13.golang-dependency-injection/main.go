package main

import (
	"fmt"
	"golang-restful-api/app/database"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/router"
	"golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := database.GetConnection()

	defer db.Close()

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	r := router.NewRouter(categoryController)

	authMiddleware := middleware.NewAuthMiddleware(r)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Listening on port 4000")
}
