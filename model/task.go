package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	User      User
	UserId    uint
	Title     string `gorm:"not null;index"`
	Content   string `gorm:"type:longText;not null"`
	Status    int    `gorm:"default 0"`
	StartTime int64
	EndTime   int64
}
