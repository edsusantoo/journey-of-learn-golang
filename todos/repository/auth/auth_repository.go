package auth

import (
	"context"
	"database/sql"
	"todos/model/domain"
)

type AuthRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
}
