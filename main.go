package main

import (
	"log"
	"net/http"

	"github.com/andiahmads/go-restfulAPI/app"
	"github.com/andiahmads/go-restfulAPI/config"
	"github.com/andiahmads/go-restfulAPI/controller"
	"github.com/andiahmads/go-restfulAPI/helper"
	"github.com/andiahmads/go-restfulAPI/middleware"
	"github.com/andiahmads/go-restfulAPI/repository"
	"github.com/andiahmads/go-restfulAPI/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	validate := validator.New()

	db := config.SetupConnection()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	//create server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Printf("server running %s", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
