package service

import (
	"context"
	"todos/model/web/auth"
)

type AuthService interface {
	Register(ctx context.Context, request auth.RegisterRequest) auth.AuthResponse
	Login(ctx context.Context, request auth.LoginRequest) (auth.AuthResponse, error)
}
