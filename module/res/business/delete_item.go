package business

import (
	resmodel "Food_Delivery3/module/res/model"
	"context"
)

type DeleteResItemStorage interface {
	FindItem(ctx context.Context, condition map[string]interface{}) (*resmodel.Restaurant, error)
	DeleteItem(ctx context.Context, condition map[string]interface{}) error
}

type deleteBiz struct {
	store DeleteResItemStorage
}

func NewDeleteToDoItemBiz