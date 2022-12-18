package test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/glenaldinlim/cake-store-service/database"
	"github.com/glenaldinlim/cake-store-service/model/entity"
	"github.com/glenaldinlim/cake-store-service/repository"
)

var cakeEntity = entity.Cake{
	Title:       "Lemon Cheesecake",
	Description: "A cheescake made of lemon",
	Rating:      7.1,
	Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
}
var cakeEntityUpdate = entity.Cake{
	Title:       "Lemon Cheesecake",
	Description: "A cheescake made of lemon",
	Rating:      8.1,
	Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
}

func truncateCakesTable(db *sql.DB) {
	db.Exec("TRUNCATE cakes")
}

func TestCakeRepoFindAll(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	tx, _ := db.Begin()

	cakeRepository := repository.NewCakeRepository()

	cakeLemon := cakeRepository.Save(context.Background(), tx, cakeEntity)
	cakeForest := cakeRepository.Save(context.Background(), tx, entity.Cake{
		Title:       "Dark Forest Cake",
		Description: "A dark forest cake made of dark chocolate",
		Rating:      9.1,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	})
	cakes := cakeRepository.FindAll(context.Background(), tx)
	tx.Commit()

	assert.Equal(t, 2, len(cakes))
	assert.Equal(t, cakeForest.Title, cakes[0].Title)
	assert.Equal(t, cakeLemon.Title, cakes[1].Title)
}

func TestCakeRepoCreate(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	tx, _ := db.Begin()

	cakeRepository := repository.NewCakeRepository()

	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	cakeResult, err := cakeRepository.FindById(context.Background(), tx, cake.Id)
	tx.Commit()

	assert.Equal(t, cakeEntity.Title, cakeResult.Title)
	assert.Equal(t, cake.Id, cakeResult.Id)
	assert.Equal(t, nil, err)
}

func TestCakeRepoUpdate(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	tx, _ := db.Begin()

	cakeRepository := repository.NewCakeRepository()

	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	cakeUpdate := cakeRepository.Update(context.Background(), tx, cakeEntity)
	cakeResult, err := cakeRepository.FindById(context.Background(), tx, cake.Id)
	tx.Commit()

	assert.Equal(t, cakeEntity.Title, cakeResult.Title)
	assert.Equal(t, cakeUpdate.Rating, cakeResult.Rating)
	assert.Equal(t, nil, err)
}

func TestCakeRepoDelete(t *testing.T) {
	db := database.InitDB()
	truncateCakesTable(db)
	tx, _ := db.Begin()

	cakeRepository := repository.NewCakeRepository()

	cake := cakeRepository.Save(context.Background(), tx, cakeEntity)
	cakeRepository.Delete(context.Background(), tx, cake)
	_, err := cakeRepository.FindById(context.Background(), tx, cake.Id)
	tx.Commit()

	assert.Error(t, err, "not found")
}
