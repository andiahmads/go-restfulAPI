package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/andiahmads/go-restfulAPI/app"
	"github.com/andiahmads/go-restfulAPI/controller"
	"github.com/andiahmads/go-restfulAPI/helper"
	"github.com/andiahmads/go-restfulAPI/middleware"
	"github.com/andiahmads/go-restfulAPI/model/domain"
	"github.com/andiahmads/go-restfulAPI/repository"
	"github.com/andiahmads/go-restfulAPI/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func SetupConnectionTest() *sql.DB {
	db, err := sql.Open("mysql", "root:endi@tcp(localhost:3306)/go_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

}

func SetupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)

}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {

	db := SetupConnectionTest()
	//delete data every running unit test
	truncateCategory(db)

	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":"Kocok"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

}

func TestCreateCategoryFailed(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test
	truncateCategory(db)

	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test
	truncateCategory(db)

	router := SetupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "melinda",
	})
	tx.Commit()
	categoryId := strconv.Itoa(category.ID)

	requestBody := strings.NewReader(`{"name":"Kocok2"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+categoryId, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test
	truncateCategory(db)

	router := SetupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "melinda",
	})
	tx.Commit()
	categoryId := strconv.Itoa(category.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+categoryId, nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}
func TestDeleteCategoryFailed(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}
func TestFindByIdCategorySuccess(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test
	truncateCategory(db)

	router := SetupRouter(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "melinda",
	})
	tx.Commit()
	categoryId := strconv.Itoa(category.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+categoryId, nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}
func TestFindByIdCategoryFailed(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}

func TestFindAllCategories(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "RAHASIA")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

}
func TestUnauthorized(t *testing.T) {
	db := SetupConnectionTest()
	//delete data every running unit test

	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
}
