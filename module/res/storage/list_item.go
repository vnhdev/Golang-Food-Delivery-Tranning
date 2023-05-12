package storage

import (
	"Food_Delivery3/common"
	"Food_Delivery3/module/res/model"
	"context"
)

func (s *mysqlStorage) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]model.Restaurant, error) {

	var result []model.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(model.RestaurantCreate{}.TableName()).Where(conditions)
	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
