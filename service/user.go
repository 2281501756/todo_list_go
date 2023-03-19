package service

import (
	"todo_list/model"
	"todo_list/package/utils"
	"todo_list/serializer"
)

type UserService struct {
	UserName string `json:"user_name" binding:"required,min=3,max=15" `
	Password string `json:"password" binding:"required,min=3,max=15" `
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(user).Where("user_name = ?", service.UserName).Count(&count)
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

func (service *UserService) Login() serializer.Response {
	var user model.User
	var count int64
	model.DB.Where("user_name = ?", service.UserName).Find(&user).Count(&count)
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "用户名不存在",
		}
	}
	flag := user.CheckPassword(service.Password)
	if !flag {
		return serializer.Response{
			Status: 403,
			Msg:    "密码不正确",
		}
	}
	token, err := utils.GeneratorToken(user.ID, user.UserName, user.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token生成失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Data: map[string]any{
			"user":  user,
			"token": token,
		},
		Msg: "登录正确",
	}
}

func GetAll() serializer.Response {
	var users []model.User
	model.DB.Find(&users)
	return serializer.Response{
		Status: 200,
		Data:   users,
	}
}
