package usermodel

import (
	"Food_Delivery3/common"
	"errors"
)

const EntityName = "User"

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Role            string        `json:"-" gorm:"column:role;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

//	func (u *UserCreate) Mask(isAdmin bool) {
//		u.GenUID(common.DbTypeUser)
//	}

var (
	ErrUserNameOrPasswordInvalid = common.NewCustomError(errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid")
	ErrEmailExisted = common.NewCustomError(errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted")
)