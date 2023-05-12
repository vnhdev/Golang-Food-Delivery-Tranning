package business

import (
	"Food_Delivery3/module/res/model"
	"context"
)

type FindSingleRestaurant interface {
	FindRestaurantById(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
}

type getResBiz struct {
	store FindSingleRestaurant
}

func NewFindRestaurantStorage(store FindSingleRestaurant) *getResBiz {
	return &getResBiz{store: store}
}

func (biz *getResBiz) GetRestaurant(ctx context.Context, id int) (*model.Restaurant, error) {
	result, err := biz.store.FindRestaurantById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
