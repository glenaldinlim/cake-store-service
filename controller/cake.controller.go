package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CakeController interface {
	Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Destroy(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
