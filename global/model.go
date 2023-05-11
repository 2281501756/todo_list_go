package global

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID       uint `gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt `json:"-" gorm:"index"`
}
