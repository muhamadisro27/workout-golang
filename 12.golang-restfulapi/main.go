package main

import (
	"fmt"
	"golang-restful-api/app/database"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.GetConnection()

	defer db.Close()

	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	r := httprouter.New()

	r.GET("/api/categories", categoryController.FindAll)
	r.POST("/api/categories", categoryController.Insert)
	r.GET("/api/categories/:categoryId", categoryController.FindById)
	r.PUT("/api/categories/:categoryId", categoryController.Update)
	r.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Listening on port 4000")
}
