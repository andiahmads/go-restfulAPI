package excepction

import (
	"net/http"

	"github.com/andiahmads/go-restfulAPI/helper"
	"github.com/andiahmads/go-restfulAPI/model/web"
	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, reqeuest *http.Request, error interface{}) {
	if notFoundError(writer, reqeuest, error) {
		return
	}

	if validationErrors(writer, reqeuest, error) {
		return
	}
	internalServerError(writer, reqeuest, error)
}

func notFoundError(writer http.ResponseWriter, reqeuest *http.Request, err interface{}) bool {
	execption, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND!",
			Data:   execption.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, reqeuest *http.Request, err interface{}) bool {
	execption, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   execption.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, reqeuest *http.Request, error interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   error,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
