package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/glenaldinlim/cake-store-service/controller"
	"github.com/glenaldinlim/cake-store-service/database"
	"github.com/glenaldinlim/cake-store-service/exception"
	"github.com/glenaldinlim/cake-store-service/model/entity"
	"github.com/glenaldinlim/cake-store-service/repository"
	"github.com/glenaldinlim/cake-store-service/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var requestBody = strings.NewReader(`{"title": "Lemon Cheesecake","description": "A cheescake made of lemon","rating": 8.1,"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}`)
var requestBodyFail = strings.NewReader(`{"description": "A cheescake made of lemon","rating": 8.1,"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}`)
var requestBodyUpdate = strings.NewReader(`{"title": "Lemon Cheesecake","description": "A cheescake made of lemon","rating": 9.9,"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}`)
var requestBodyUpdateFail = strings.NewReader(`{"description": "A cheescake made of lemon","rating": 8.1,"image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}`)

func SetupRouter(db *sql.DB) *httprouter.Router {
	validate := validator.New()
	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)
	cakeController := controller.NewCakeController(cakeService)

	router := httprouter.New()

	router.GET("/api/cakes", cakeController.Index)
	router.POST("/api/cakes", cakeController.Store)
	router.GET("/api/cakes/:id", cakeController.Show)
	router.PATCH("/api/cakes/:id", cakeController.Update)
	router.DELETE("/api/cakes/:id", cakeController.Destroy)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func TestCakeControllerCreateSuccess(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8090/api/cakes", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
	assert.Equal(t, "Lemon Cheesecake", resBody["data"].(map[string]interface{})["title"])
}

func TestCakeControllerCreateFailed(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8090/api/cakes", requestBodyFail)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 400, int(resBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", resBody["status"])
}

func TestCakeControllerUpdateSuccess(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	tx.Commit()

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8090/api/cakes/"+strconv.Itoa(int(cake.Id)), requestBodyUpdate)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
	assert.Equal(t, int(cake.Id), int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 9.9, resBody["data"].(map[string]interface{})["rating"])
}

func TestCakeControllerUpdateFailed(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	tx.Commit()

	request := httptest.NewRequest(http.MethodPatch, "http://localhost:8090/api/cakes/"+strconv.Itoa(int(cake.Id)), requestBodyUpdateFail)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 400, int(resBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", resBody["status"])
}

func TestCakeControllerGetByIdSuccess(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8090/api/cakes/"+strconv.Itoa(int(cake.Id)), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
	assert.Equal(t, int(cake.Id), int(resBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, cake.Title, resBody["data"].(map[string]interface{})["title"])
}

func TestCakeControllerGetByIdFailed(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8090/api/cakes/100", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 404, int(resBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", resBody["status"])
}

func TestCakeControllerDeleteSuccess(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	tx.Commit()

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8090/api/cakes/"+strconv.Itoa(int(cake.Id)), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])
}

func TestCakeControllerDeleteFailed(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8090/api/cakes/100", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 404, int(resBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", resBody["status"])
}

func TestCakeControllerGetAll(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	router := SetupRouter(db)

	tx, _ := db.Begin()
	cakeRepository := repository.NewCakeRepository()
	cakeLemon := cakeRepository.Save(context.Background(), tx, cakeEntity)
	cakeForest := cakeRepository.Save(context.Background(), tx, entity.Cake{
		Title:       "Dark Forest Cake",
		Description: "A dark forest cake made of dark chocolate",
		Rating:      9.1,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	})
	tx.Commit()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8090/api/cakes", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var resBody map[string]interface{}
	json.Unmarshal(body, &resBody)

	assert.Equal(t, 200, int(resBody["code"].(float64)))
	assert.Equal(t, "OK", resBody["status"])

	var cakes = resBody["data"].([]interface{})
	cakeResForest := cakes[0].(map[string]interface{})
	cakeResLemon := cakes[1].(map[string]interface{})

	assert.Equal(t, int(cakeForest.Id), int(cakeResForest["id"].(float64)))
	assert.Equal(t, cakeForest.Title, cakeResForest["title"])

	assert.Equal(t, int(cakeLemon.Id), int(cakeResLemon["id"].(float64)))
	assert.Equal(t, cakeLemon.Title, cakeResLemon["title"])
}
