package app

import (
	"net/http"
	auth_controller "todos/controller/auth"
	todo_controller "todos/controller/todo"
	"todos/exception"
	"todos/helper"
	"todos/model/web"

	"github.com/julienschmidt/httprouter"
)

func NewRoute(authController auth_controller.AuthController, todoController todo_controller.TodoController) *httprouter.Router {
	router := httprouter.New()

	//auth
	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/register", authController.Register)
	router.POST("/api/todo", todoController.Create)
	router.PATCH("/api/todo/:id", todoController.Update)
	router.DELETE("/api/todo/:id", todoController.Delete)
	router.GET("/api/todo/:id", todoController.FindById)
	router.GET("/api/todos", todoController.FindAll)

	router.PanicHandler = exception.ErrorHandler
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Page not found",
			Data:   nil,
		}

		helper.WriteToResponseBody(w, webResponse)
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		webResponse := web.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "Method not allowed",
			Data:   nil,
		}

		helper.WriteToResponseBody(w, webResponse)
	})
	return router
}
