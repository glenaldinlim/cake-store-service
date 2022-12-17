package utils

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(location string, request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(location, err)
}

func WriteResponseBody(location string, writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(location, err)
}
