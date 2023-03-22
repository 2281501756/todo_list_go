package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	User      User   `json:"user"`
	UserId    uint   `json:"userId"`
	Title     string `json:"title" gorm:"not null;index"`
	Content   string `json:"content" gorm:"type:longText;not null"`
	Status    int    `json:"status" gorm:"default 0"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}
