package todo

import (
	"net/http"
	"strconv"
	"todos/helper"
	"todos/model/web"
	todo_web "todos/model/web/todo"
	todo_service "todos/service/todo"

	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService todo_service.TodoService
}

func NewTodoController(todoService todo_service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

func (controller *TodoControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	createRequest := todo_web.TodoCreateRequest{}
	helper.ReadFromRequestBody(r, &createRequest)

	todoResponse := controller.TodoService.Create(r.Context(), createRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	updateRequest := todo_web.TodoUpdateRequest{}
	helper.ReadFromRequestBody(r, &updateRequest)

	todoId := params.ByName("id")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	updateRequest.ID = id

	todoResponse := controller.TodoService.Update(r.Context(), updateRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	controller.TodoService.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoResponse := controller.TodoService.FindById(r.Context(), id)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *TodoControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todoResponse := controller.TodoService.FindAll(r.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
