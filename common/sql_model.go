package common

import "time"

type SQLModel struct {
	FakeId   *UID       `json:"id" gorm:"-"`
	Id       int        `json:"-" gorm:"column:id;"`
	Status   int        `json:"status" gorm:"column:status;"`
	CreateAt *time.Time `json:"create_at,omitempty" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
