package main

import (
	"fmt"
	"net/http"
	"todos/app"
	"todos/helper"

	"github.com/go-playground/validator/v10"

	controller_auth "todos/controller/auth"
	repository_auth "todos/repository/auth"
	service_auth "todos/service/auth"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	authRepo := repository_auth.NewAuthRepository()
	authService := service_auth.NewAuthService(authRepo, db, validate)
	authController := controller_auth.NewAuthController(authService)

	router := app.NewRoute(authController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Starting serve in localhost:3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
