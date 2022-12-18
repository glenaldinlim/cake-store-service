package test

import (
	"context"
	"testing"

	"github.com/glenaldinlim/cake-store-service/database"
	"github.com/glenaldinlim/cake-store-service/model/web"
	"github.com/glenaldinlim/cake-store-service/repository"
	"github.com/glenaldinlim/cake-store-service/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var cakeReq = web.CakeRequest{
	Title:       "Lemon Cheesecake",
	Description: "A cheescake made of lemon",
	Rating:      7.1,
	Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
}

func TestCakeSrvFindAll(t *testing.T) {
	db := database.InitDB()
	validate := validator.New()
	truncateCakesTable(db)

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)

	cakeLemon := cakeService.Create(context.Background(), cakeReq)
	cakeForest := cakeService.Create(context.Background(), web.CakeRequest{
		Title:       "Dark Forest Cake",
		Description: "A dark forest cake made of dark chocolate",
		Rating:      9.1,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	})
	cakeRes := cakeService.FindAll(context.Background())

	assert.Equal(t, 2, len(cakeRes))
	assert.Equal(t, cakeForest.Title, cakeRes[0].Title)
	assert.Equal(t, cakeLemon.Title, cakeRes[1].Title)
}

func TestCakeSrvCreate(t *testing.T) {
	db := database.InitDB()
	validate := validator.New()
	truncateCakesTable(db)

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)

	cake := cakeService.Create(context.Background(), cakeReq)
	cakeRes := cakeService.FindById(context.Background(), cake.Id)

	assert.Equal(t, cakeReq.Title, cakeRes.Title)
	assert.Equal(t, cakeReq.Rating, cakeRes.Rating)
}

func TestCakeSrvUpdate(t *testing.T) {
	db := database.InitDB()
	validate := validator.New()
	truncateCakesTable(db)

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)

	cake := cakeService.Create(context.Background(), cakeReq)
	cakeUpdate := cakeService.Update(context.Background(), web.CakeRequest{
		Title:       "Lemon Cheesecake",
		Description: "A cheescake made of lemon",
		Rating:      8.1,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}, cake.Id)
	cakeRes := cakeService.FindById(context.Background(), cake.Id)

	assert.Equal(t, cakeReq.Title, cakeRes.Title)
	assert.Equal(t, cakeUpdate.Rating, cakeRes.Rating)
}

func TestCakeSrvDelete(t *testing.T) {
	db := database.InitDB()
	validate := validator.New()
	truncateCakesTable(db)

	cakeRepository := repository.NewCakeRepository()
	cakeService := service.NewCakeService(cakeRepository, db, validate)

	cake := cakeService.Create(context.Background(), cakeReq)
	cakeService.Delete(context.Background(), cake.Id)
}
