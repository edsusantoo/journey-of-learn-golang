package todo

import (
	"context"
	"database/sql"
	"errors"
	"todos/helper"
	"todos/model/domain"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repo *TodoRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "insert into todos (title,body,date_time) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, todo.Title, todo.Body, todo.DateTime)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	todo.Id = int(id)

	return todo
}

func (repo *TodoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL_UPDATE := "update todos set title = ?, body = ?, date_time = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL_UPDATE, todo.Title, todo.Body, todo.DateTime, todo.Id)
	helper.PanicIfError(err)

	return todo

}

func (repo *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) {
	SQL_DELETE := "delete from todos where id = ?"
	_, err := tx.ExecContext(ctx, SQL_DELETE, todo.Id)
	helper.PanicIfError(err)
}

func (repo *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Todo, error) {
	SQL := "select * from todos where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.Todo{}

	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.DateTime, &todo.CreatedAt, &todo.UpdatedAt)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}
}

func (repo *TodoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo {
	SQL := "select * from todos"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.Todo
	for rows.Next() {
		todo := domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Body, &todo.DateTime, &todo.CreatedAt, &todo.UpdatedAt)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	return todos
}
