package exception

import (
	"fmt"
	"net/http"

	"github.com/glenaldinlim/cake-store-service/model/web"
	"github.com/glenaldinlim/cake-store-service/utils"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	res := web.WebResponse{
		Code:     http.StatusInternalServerError,
		Status:   "INTERNAL SERVER ERROR",
		Messsage: fmt.Sprintf("%v", err),
	}

	utils.WriteResponseBody("Exception: InternalServerError", writer, res)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		res := web.WebResponse{
			Code:     http.StatusNotFound,
			Status:   "NOT FOUND",
			Messsage: exception.Error,
		}

		utils.WriteResponseBody("Exception: InternalServerError", writer, res)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		res := web.WebResponse{
			Code:     http.StatusBadRequest,
			Status:   "BAD REQUEST",
			Messsage: exception.Error(),
		}

		utils.WriteResponseBody("Exception: InternalServerError", writer, res)
		return true
	} else {
		return false
	}
}
