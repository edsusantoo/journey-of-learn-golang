package auth

import (
	"net/http"
	"todos/helper"
	"todos/jwt"
	"todos/model/web"
	web_auth "todos/model/web/auth"
	service_auth "todos/service/auth"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	AuthService service_auth.AuthService
}

func NewAuthController(service service_auth.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: service,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web_auth.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse, err := controller.AuthService.Login(request.Context(), loginRequest)

	if err != nil {
		helper.WriteToErrorBody(writer, 400, err)
	} else {
		token, err := jwt.CreateToken(loginResponse.Username)
		helper.WriteToErrorBody(writer, 400, err)

		helper.WriteToResponseBody(writer, web.LoginResponse{
			Code:   200,
			Status: "success",
			Data:   loginResponse,
			Token:  token,
		})
	}
}

func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerRequest := web_auth.RegisterRequest{}
	helper.ReadFromRequestBody(request, &registerRequest)

	registerResponse := controller.AuthService.Register(request.Context(), registerRequest)

	helper.WriteToResponseBody(writer, web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   registerResponse,
	})
}
