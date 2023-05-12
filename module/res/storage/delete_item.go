package storage

import (
	"Food_Delivery3/module/res/model"
	"context"
)

func (s *mysqlStorage) SoftDeleteData(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}

	return nil
}
