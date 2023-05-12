package storage

import (
	"Food_Delivery3/common"
	"Food_Delivery3/module/res/model"
	"context"
	"gorm.io/gorm"
)

func (s *mysqlStorage) FindRestaurantById(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*model.Restaurant, error) {
	var data model.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//data not found
			return nil, common.RecordNotFound
		}
		return nil, err
	}
	return &data, nil
}
