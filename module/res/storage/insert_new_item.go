package storage

import (
	"Food_Delivery3/module/res/model"
	"context"
)

func (s *mysqlStorage) CreateRes(ctx context.Context, data *model.Restaurant) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
