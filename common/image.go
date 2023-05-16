package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB %s", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return nil
	}
	*j = img
	return nil
}
