package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-restful-api/controller"
	"golang-restful-api/entity/domain"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/router"
	"golang-restful-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	dsn := "root:123456@unix(/home/user/workout-golang/12.golang-restfulapi/mysql/data/mysql.sock)/belajar_golang_restful_api_test?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	helper.PanicIfError(err)

	return db
}

func TestDBConnection(t *testing.T) {
	db := setupTestDB()

	defer db.Close()
}

func setupRouter(DB *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, DB, validate)
	categoryController := controller.NewCategoryController(categoryService)

	r := router.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(r)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	r := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Isro"}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:4000/api/categories", requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Isro", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailedRequestInvalid(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	r := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:4000/api/categories", requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Isro Edit"}`)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Isro Edit", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailedRequestNotFound(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Isro Edit"}`)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/categories/"+strconv.Itoa((category.Id+1)), requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestUpdateCategoryFailedRequestInvalid(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)

	req := httptest.NewRequest(http.MethodPut, "http://localhost:4000/api/categories/"+strconv.Itoa(category.Id), requestBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:4000/api/categories/"+strconv.Itoa(category.Id), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailedRequestNotFound(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	req := httptest.NewRequest(http.MethodDelete, "http://localhost:4000/api/categories/"+strconv.Itoa((category.Id+1)), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestFindByIdCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories/"+strconv.Itoa(category.Id), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Isro", responseBody["data"].(map[string]interface{})["name"])
}

func TestFindByIdCategoryFailedRequestNotFound(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)

	tx, _ := db.Begin()

	categoryRepository := repository.NewCategoryRepository()

	category := categoryRepository.Insert(context.Background(), tx, domain.Category{
		Name: "Isro",
	})

	tx.Commit()

	r := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories/"+strconv.Itoa((category.Id+1)), nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)
	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestFindAllCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	r := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", "RAHASIA")

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)

	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	fmt.Println(responseBody)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	r := setupRouter(db)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:4000/api/categories", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(rec, req)

	res := rec.Result()

	bytes, err := io.ReadAll(res.Body)
	helper.PanicIfError(err)

	var responseBody map[string]interface{}

	err = json.Unmarshal(bytes, &responseBody)
	helper.PanicIfError(err)

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
