package model

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"todo_list/cache"
)

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

func (this *Task) GetView() uint64 {
	countStr, _ := cache.RDB.Get(context.Background(), cache.GetTaskViewKey(this.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (this *Task) ViewAdd() {
	cache.RDB.Incr(context.Background(), cache.GetTaskViewKey(this.ID))
	cache.RDB.ZIncrBy(context.Background(), cache.TaskRankKey, 1, strconv.Itoa(int(this.ID)))
}
