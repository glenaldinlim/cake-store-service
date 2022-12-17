package controller

import (
	"net/http"
	"strconv"

	"github.com/glenaldinlim/cake-store-service/model/web"
	"github.com/glenaldinlim/cake-store-service/service"
	"github.com/glenaldinlim/cake-store-service/utils"
	"github.com/julienschmidt/httprouter"
)

type CakeControllerImpl struct {
	CakeService service.CakeService
}

func NewCakeController(cakeService service.CakeService) CakeController {
	return &CakeControllerImpl{
		CakeService: cakeService,
	}
}

func (controller *CakeControllerImpl) Index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakes := controller.CakeService.FindAll(request.Context())
	res := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cakes,
	}

	utils.WriteResponseBody("CakeController.Index", writer, res)
}

func (controller *CakeControllerImpl) Store(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeRequest := web.CakeRequest{}
	utils.ReadRequestBody("CakeController.Store", request, &cakeRequest)

	cake := controller.CakeService.Create(request.Context(), cakeRequest)
	res := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cake,
	}

	utils.WriteResponseBody("CakeController.Store", writer, res)
}

func (controller *CakeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeRequest := web.CakeRequest{}
	utils.ReadRequestBody("CakeController.Update", request, &cakeRequest)

	cakeId := params.ByName("id")
	id, err := strconv.ParseInt(cakeId, 10, 64)
	if err != nil {
		utils.Logger().WithField("cakeId", cakeId).Errorf("[Parse] CakeController.Update: %s", err.Error())
		panic(err)
	}

	cake := controller.CakeService.Update(request.Context(), cakeRequest, id)
	res := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cake,
	}

	utils.WriteResponseBody("CakeController.Update", writer, res)
}

func (controller *CakeControllerImpl) Show(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeId := params.ByName("id")
	id, err := strconv.ParseInt(cakeId, 10, 64)
	if err != nil {
		utils.Logger().WithField("cakeId", cakeId).Errorf("[Parse] CakeController.Show: %s", err.Error())
		panic(err)
	}

	cake := controller.CakeService.FindById(request.Context(), id)
	res := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cake,
	}

	utils.WriteResponseBody("CakeController.Show", writer, res)
}

func (controller *CakeControllerImpl) Destroy(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cakeId := params.ByName("id")
	id, err := strconv.ParseInt(cakeId, 10, 64)
	if err != nil {
		utils.Logger().WithField("cakeId", cakeId).Errorf("[Parse] CakeController.Destroy: %s", err.Error())
		panic(err)
	}

	controller.CakeService.Delete(request.Context(), id)
	res := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	utils.WriteResponseBody("CakeController.Destroy", writer, res)
}
