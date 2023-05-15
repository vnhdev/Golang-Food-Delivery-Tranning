package common

import (
	"database/sql/driver"
	"encoding/json"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name" gorm:"-"`
	Extension string `json:"extension" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

// De implement struct xuong Database thi ta dung ham Value()
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
