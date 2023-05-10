package storage

import (
	"Food_Delivery3/module/res/model"
	"context"
)

func (s *mysqlStorage) ListRes(ctx context.Context, data *model.Restaurant) error {
	if err := s.db.Find(data).Error; err != nil {
		return err
	}
	return nil
}
