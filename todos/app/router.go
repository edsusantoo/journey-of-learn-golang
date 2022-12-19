package app

import (
	"net/http"
	controller_auth "todos/controller/auth"
	"todos/exception"
	"todos/helper"
	"todos/model/web"

	"github.com/julienschmidt/httprouter"
)

func NewRoute(authController controller_auth.AuthController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/register", authController.Register)
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
