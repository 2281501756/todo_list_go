package model

import "todo_list/global"

type UserModel struct {
	global.MODEL
	UserName string `json:"username" gorm:"comment:用户名"`
	Password string `json:"password" gorm:"comment:密码"`
	Name     string `json:"name" gorm:"comment:昵称"`
}
