package model

import (
	"Food_Delivery3/common"
	"errors"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	// ,inline o day la de cac tag trong common.SQLModel co cung cap voi cac tag cua Restaurant
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name"`
	Addr            string        `json:"addr" gorm:"column:addr"`
	Logo            *common.Image `json:"logo" gorm:"logo;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string       `json:"name" gorm:"column:name"`
	Addr *string       `json:"addr" gorm:"column:addr"`
	Logo *common.Image `json:"logo" gorm:"logo;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Name string        `json:"name" gorm:"column:name"`
	Addr string        `json:"addr" gorm:"column:addr"`
	Logo *common.Image `json:"logo" gorm:"logo;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}
	return nil
}
