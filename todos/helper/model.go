package helper

import (
	"todos/model/domain"
	"todos/model/web/auth"
	"todos/model/web/todo"
)

func ToAuthResponse(domain domain.User) auth.AuthResponse {
	return auth.AuthResponse{
		Id:        domain.Id,
		Username:  domain.Username,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Active:    domain.Active,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToTodoResponse(domain domain.Todo) todo.TodoResponse {
	return todo.TodoResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Body:      domain.Body,
		DateTime:  domain.DateTime,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ToTodoResponses(domain []domain.Todo) []todo.TodoResponse {
	var todoResponses []todo.TodoResponse
	for _, todo := range domain {
		todoResponses = append(todoResponses, ToTodoResponse(todo))
	}

	return todoResponses
}
