package main

import (
	"net/http"

	"github.com/glenaldinlim/cake-store-service/controller"
	"github.com/glenaldinlim/cake-store-service/database"
	"github.com/glenaldinlim/cake-store-service/exception"
	"github.com/glenaldinlim/cake-store-service/repository"
	"github.com/glenaldinlim/cake-store-service/service"
	"github.com/glenaldinlim/cake-store-service/utils"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.InitDB()
	validate := validator.New()

	cakeRepo := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepo, db, validate)
	cakeController := controller.NewCakeController(cakeService)

	router := httprouter.New()

	router.GET("/api/cakes", cakeController.Index)
	router.POST("/api/cakes", cakeController.Store)
	router.GET("/api/cakes/:id", cakeController.Show)
	router.PATCH("/api/cakes/:id", cakeController.Update)
	router.DELETE("/api/cakes/:id", cakeController.Destroy)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "0.0.0.0:8090",
		Handler: router,
	}

	utils.Logger().Infof("service running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		utils.Logger().Errorf("failed to running service: %s", err.Error())
		panic(err)
	}
}
