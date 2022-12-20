package jwt

import (
	"time"

	"todos/helper"

	"github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = "TODOS"
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret-key-Rootkit")
var expiredToken = helper.GetDateTimeNowDate().Add(30 * time.Minute)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Token struct {
	Data      string `json:"data"`
	ExpiredAt string `json:"expired_at"`
}

func CreateToken(username string) (Token, error) {
	claims := &MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredToken),
			Issuer:    APPLICATION_NAME,
		},
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)

	return Token{signedToken, expiredToken.Format("2006-01-02 15:04:05")}, err
}
