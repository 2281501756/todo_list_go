package global

import (
	"gorm.io/gorm"
	"todo_list/config"
)

var (
	Config config.Config
	DB     *gorm.DB
)
