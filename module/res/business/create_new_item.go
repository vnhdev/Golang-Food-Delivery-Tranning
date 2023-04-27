package business

import (
	"Food_Delivery3/module/res/model"
	"context"
	"errors"
)

type CreateResItemStorage interface {
	CreateRes(ctx context.Context, data *model.Restaurant) error
}

type createBiz struct {
	store CreateResItemStorage
}

func NewCreateResItemBiz(store CreateResItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewRes(ctx context.Context, data *model.Restaurant) error {
	if data.Name == "" {
		return errors.New("Name can't be blank")
	}

	if err := biz.store.CreateRes(ctx, data); err != nil {
		return err
	}
	return nil
}
