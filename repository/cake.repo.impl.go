package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/glenaldinlim/cake-store-service/model/entity"
	"github.com/glenaldinlim/cake-store-service/utils"
)

type CakeRepositoryImpl struct {
}

func NewCakeRepository() CakeRepository {
	return &CakeRepositoryImpl{}
}

func (repo *CakeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Cake {
	SQL := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes ORDER BY rating DESC, title ASC"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		utils.Logger().Errorf("[DB Exception] CakeRepository.FindAll: %s", err.Error())
		panic(err)
	}
	defer rows.Close()

	var cakes []entity.Cake
	for rows.Next() {
		cake := entity.Cake{}
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			utils.Logger().Errorf("[DB Exception] CakeRepository.FindAll: %s", err.Error())
			panic(err)
		}
		cakes = append(cakes, cake)
	}
	return cakes
}

func (repo *CakeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, cakeId int64) (entity.Cake, error) {
	SQL := "SELECT id, title, description, rating, image, created_at, updated_at FROM cakes WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, cakeId)
	if err != nil {
		utils.Logger().Errorf("[DB Exception] CakeRepository.FindById: %s", err.Error())
		panic(err)
	}
	defer rows.Close()

	cake := entity.Cake{}
	if rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			utils.Logger().Errorf("[DB Exception] CakeRepository.FindById: %s", err.Error())
			panic(err)
		}
		return cake, nil
	} else {
		utils.Logger().WithField("cakeId", cakeId).Warn("cake not found")
		return cake, errors.New("cake not found")
	}
}

func (repo *CakeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, cake entity.Cake) entity.Cake {
	SQL := "INSERT INTO cakes(title, description, rating, image) VALUES (?, ?, ?, ?)"
	res, err := tx.ExecContext(ctx, SQL, cake.Title, cake.Description, cake.Rating, cake.Image)
	if err != nil {
		utils.Logger().Errorf("[DB Exception] CakeRepository.Save: %s", err.Error())
		panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		utils.Logger().Errorf("[DB Exception] CakeRepository.Save: %s", err.Error())
		panic(err)
	}

	cake.Id = id
	return cake
}

func (repo *CakeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, cake entity.Cake) entity.Cake {
	SQL := "UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, cake.Title, cake.Description, cake.Rating, cake.Image, cake.UpdatedAt, cake.Id)
	if err != nil {
		utils.Logger().Errorf("[DB Exception] CakeRepository.Update: %s", err.Error())
		panic(err)
	}

	return cake
}

func (repo *CakeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, cake entity.Cake) {
	SQL := "DELETE FROM cakes WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, cake.Id)
	if err != nil {
		utils.Logger().Errorf("[DB Exception] CakeRepository.Delete: %s", err.Error())
		panic(err)
	}
}
