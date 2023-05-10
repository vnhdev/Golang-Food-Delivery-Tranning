package business

import (
	"Food_Delivery3/module/res/model"
	"context"
)

type CreateResItemStorage interface {
	CreateRes(ctx context.Context, data *model.RestaurantCreate) error
}

type createBiz struct {
	store CreateResItemStorage
}

func NewCreateResItemBiz(store CreateResItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewRestaurant(ctx context.Context, data *model.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.CreateRes(ctx, data); err != nil {
		return err
	}
	return nil
}
