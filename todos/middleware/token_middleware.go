package middleware

import (
	"errors"
	"net/http"
	"strings"
	"todos/helper"
	my_jwt "todos/jwt"
	"todos/model/web"

	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/auth/login" || r.URL.Path == "/api/auth/register" {
		middleware.Handler.ServeHTTP(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	authorizationHeader := r.Header.Get("Authorization")

	if len(authorizationHeader) == 0 {
		helper.WriteToResponseBody(w, web.MessageResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Cannot find Authorization",
		})
		return
	}

	if !strings.Contains(authorizationHeader, "Bearer") {
		helper.WriteToResponseBody(w, web.MessageResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Format token invalid",
		})
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return my_jwt.JWT_SIGNATURE_KEY, nil
	})

	if token.Valid {
		middleware.Handler.ServeHTTP(w, r)
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		helper.WriteToResponseBody(w, web.MessageResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "That's not even a token",
		})
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		helper.WriteToResponseBody(w, web.MessageResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Token is expired",
		})
	} else {
		helper.WriteToResponseBody(w, web.MessageResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Couldn't handle this token",
		})
	}

}
