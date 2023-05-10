package business

//import (
//	"Food_Delivery3/module/res/model"
//	"context"
//	"errors"
//)
//
//type CreateListResStorage interface {
//	listRes(ctx context.Context, data *model.Restaurant) error
//}
//
//type listBiz struct {
//	store CreateListResStorage
//}
//
//func NewListResItemBiz(store CreateListResStorage) *listBiz {
//	return &listBiz{store: store}
//}
//
//func (biz *createBiz) ListDataByCondition(ctx context.Context, filter *model.Restaurant) error {
//	if data.Name == "" {
//		return errors.New("Name can't be blank")
//	}
//
//	if err := biz.store.CreateRes(ctx, data); err != nil {
//		return err
//	}
//	return nil
//}
