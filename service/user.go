package service

import (
	"todo_list/model"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" `
	Password string `form:"password" json:"password" binding:"required,min=3,max=15" `
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(user).Where("user_name = ?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "用户名已存在",
		}
	}
	user.UserName = service.UserName
	if user.SetPassword(service.Password) != nil {
		return serializer.Response{Status: 400, Msg: "密码不正确"}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{Status: 500, Msg: "数据库错误"}
	}
	return serializer.Response{Status: 200, Data: service, Msg: "注册成功"}
}
