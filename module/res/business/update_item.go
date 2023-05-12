package business

import (
	"Food_Delivery3/module/res/model"
	"context"
	"errors"
)

type UpdateRestaurantStorage interface {
	UpdateData(
		ctx context.Context,
		id int,
		data *model.RestaurantUpdate) error
	FindRestaurantById(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStorage
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStorage) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *model.RestaurantUpdate) error {
	oldData, err := biz.store.FindRestaurantById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
