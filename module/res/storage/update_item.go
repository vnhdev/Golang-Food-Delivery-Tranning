package storage

import (
	"Food_Delivery3/module/res/model"
	"context"
)

func (s *mysqlStorage) UpdateData(
	ctx context.Context,
	id int,
	data *model.RestaurantUpdate,
) error {
	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
