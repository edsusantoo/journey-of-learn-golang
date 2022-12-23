package todo

import (
	"context"
	"database/sql"
	"todos/exception"
	"todos/helper"
	"todos/model/domain"
	"todos/model/web/todo"
	todo_repo "todos/repository/todo"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	TodoRepository todo_repo.TodoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(repo todo_repo.TodoRepository, DB *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: repo,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request todo.TodoCreateRequest) todo.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo := domain.Todo{
		Title:    request.Title,
		Body:     request.Body,
		DateTime: request.DateTime,
	}

	todo = service.TodoRepository.Create(ctx, tx, todo)

	return helper.ToTodoResponse(todo)

}

func (service *TodoServiceImpl) Update(ctx context.Context, request todo.TodoUpdateRequest) todo.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	mTodo, err := service.TodoRepository.FindById(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	mTodo.Title = request.Title
	mTodo.Body = request.Body
	mTodo.DateTime = request.DateTime

	mTodo = service.TodoRepository.Update(ctx, tx, mTodo)

	return helper.ToTodoResponse(mTodo)

}

func (service *TodoServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoRepository.Delete(ctx, tx, todo)
}

func (service *TodoServiceImpl) FindById(ctx context.Context, id int) todo.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) FindAll(ctx context.Context) []todo.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todos := service.TodoRepository.FindAll(ctx, tx)
	return helper.ToTodoResponses(todos)
}
