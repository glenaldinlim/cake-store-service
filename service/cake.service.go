package service

import (
	"context"

	"github.com/glenaldinlim/cake-store-service/model/entity"
	"github.com/glenaldinlim/cake-store-service/model/web"
)

type CakeService interface {
	FindAll(ctx context.Context) []entity.Cake
	FindById(ctx context.Context, cakeId int64) entity.Cake
	Create(ctx context.Context, request web.CakeRequest) entity.Cake
	Update(ctx context.Context, request web.CakeRequest, cakeId int64) entity.Cake
	Delete(ctx context.Context, cakeId int64)
}
