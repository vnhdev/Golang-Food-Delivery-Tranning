package storage

import (
	"Food_Delivery3/module/res/model"
	"context"
)

func (s *mysqlStorage) CreateRes(ctx context.Context, data *model.RestaurantCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
