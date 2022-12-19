package helper

import (
	"encoding/json"
	"net/http"
	"todos/model/web"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteToErrorBody(writer http.ResponseWriter, code int, err error) {
	if err != nil {
		response := web.WebResponse{
			Code:   code,
			Status: err.Error(),
			Data:   nil,
		}
		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(response)
		PanicIfError(err)
	}
}
