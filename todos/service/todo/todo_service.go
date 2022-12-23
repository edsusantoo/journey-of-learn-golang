package todo

import (
	"context"
	"todos/model/web/todo"
)

type TodoService interface {
	Create(ctx context.Context, request todo.TodoCreateRequest) todo.TodoResponse
	Update(ctx context.Context, request todo.TodoUpdateRequest) todo.TodoResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) todo.TodoResponse
	FindAll(ctx context.Context) []todo.TodoResponse
}
