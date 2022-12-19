package auth

import (
	"context"
	"database/sql"
	"errors"
	"todos/helper"
	"todos/model/domain"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (username,first_name,last_name,password)values(?,?,?,?)"
	result, err := tx.ExecContext(ctx,
		SQL,
		user.Username,
		user.FirstName,
		user.LastName,
		helper.HashPassword(user.Password),
	)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)

	return user

}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {

	SQL := "SELECT * FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username)
	helper.PanicIfError(err)
	defer rows.Close()

	dUser := domain.User{}

	if rows.Next() {
		err := rows.Scan(
			&dUser.Id,
			&dUser.Username,
			&dUser.Password,
			&dUser.FirstName,
			&dUser.LastName,
			&dUser.Active,
			&dUser.CreatedAt,
			&dUser.UpdatedAt)
		helper.PanicIfError(err)

		//check password
		if helper.CheckPasswordHash(user.Password, dUser.Password) {
			return dUser, nil
		} else {
			return dUser, errors.New("password not match")
		}

	} else {
		return dUser, errors.New("user not found")
	}
}
