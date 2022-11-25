package service

import (
	"context"
	"database/sql"
	"todos/helper"
	"todos/model/domain"
	authWeb "todos/model/web/auth"
	authRepo "todos/repository/auth"

	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepository authRepo.AuthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *AuthServiceImpl) Register(ctx context.Context, request authWeb.RegisterRequest) authWeb.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
	}

	user = service.AuthRepository.Register(ctx, tx, user)

	return helper.ToAuthResponse(user)
}

func (service *AuthServiceImpl) Login(ctx context.Context, request authWeb.LoginRequest) (authWeb.AuthResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err = service.AuthRepository.Login(ctx, tx, user)

	return helper.ToAuthResponse(user), err
}
