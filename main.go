package main

import (
	"belajar-golang-restapi/helper"
	"belajar-golang-restapi/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func NewValidation() *validator.Validate {
	return validator.New()
}

func main() {
	//db := app.NewDB()
	//validate := validator.New()
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//
	//router := app.NewRouter(categoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)

	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
