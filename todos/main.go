package main

import (
	"fmt"
	"net/http"
	"todos/app"
	"todos/helper"
	"todos/middleware"

	"github.com/go-playground/validator/v10"

	auth_controller "todos/controller/auth"
	auth_repository "todos/repository/auth"
	auth_service "todos/service/auth"

	todo_controller "todos/controller/todo"
	todo_repository "todos/repository/todo"
	todo_service "todos/service/todo"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	authRepo := auth_repository.NewAuthRepository()
	authService := auth_service.NewAuthService(authRepo, db, validate)
	authController := auth_controller.NewAuthController(authService)

	todoRepo := todo_repository.NewTodoRepository()
	todoService := todo_service.NewTodoService(todoRepo, db, validate)
	todoController := todo_controller.NewTodoController(todoService)

	router := app.NewRoute(authController, todoController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Starting serve in localhost:3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
