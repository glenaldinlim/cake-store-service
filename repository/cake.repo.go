package repository

import (
	"context"
	"database/sql"

	"github.com/glenaldinlim/cake-store-service/model/entity"
)

type CakeRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Cake
	FindById(ctx context.Context, tx *sql.Tx, cakeId int64) (entity.Cake, error)
	Save(ctx context.Context, tx *sql.Tx, cake entity.Cake) entity.Cake
	Update(ctx context.Context, tx *sql.Tx, cake entity.Cake) entity.Cake
	Delete(ctx context.Context, tx *sql.Tx, cake entity.Cake)
}
