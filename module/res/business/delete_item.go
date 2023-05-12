package business

import (
	"Food_Delivery3/module/res/model"
	"context"
)

type DeleteResItemStorage interface {
	FindRestaurantById(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
	SoftDeleteData(ctx context.Context, id int) error
}

type deleteBiz struct {
	store DeleteResItemStorage
}

func NewDeleteResItemBiz(store DeleteResItemStorage) *deleteBiz {
	return &deleteBiz{store: store}
}

func (biz *deleteBiz) DeleteRestaurant(ctx context.Context, id int) error {
	// Step 1: Find item by condition
	_, err := biz.store.FindRestaurantById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil
	}

	// Step 2: call storage to delete item
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
