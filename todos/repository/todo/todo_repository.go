package todo

import (
	"context"
	"database/sql"
	"todos/model/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todo, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo
}
