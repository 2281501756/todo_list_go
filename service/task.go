package service

import (
	"fmt"
	"todo_list/model"
	"todo_list/serializer"
)

type TaskCreateService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Status  int    `json:"status" form:"status"`
}

type TaskUpdateService struct {
	Id      uint   `json:"id" form:"id" binding:"required"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

func GetALlTask() serializer.Response {
	var task []model.Task
	model.DB.Model(model.Task{}).Preload("User").Find(&task)
	return serializer.Response{
		Status: 200,
		Data:   task,
		Msg:    "查询成功",
	}
}
func DeleteTask(id uint) serializer.Response {
	var task model.Task
	var count int64
	model.DB.Find(&task, id).Count(&count)
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "任务不存在",
		}
	}
	model.DB.Delete(&task)
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
		Data:   task,
	}
}

func (service *TaskCreateService) Create(id uint) serializer.Response {
	var task model.Task
	var user model.User
	model.DB.Find(&user, id)
	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status
	task.User = user
	task.UserId = id
	err := model.DB.Create(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   task,
		Msg:    "创建成功",
	}
}

func (service *TaskUpdateService) Update() serializer.Response {
	var task model.Task
	var count int64
	err := model.DB.Preload("User").Find(&task, service.Id).Count(&count).Updates(model.Task{Title: service.Title, Content: service.Content, Status: service.Status}).Error
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Msg:    fmt.Sprintf("更新失败,不存在id为%d的这个笔记", service.Id),
		}
	}
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "更新失败",
			Data:   task,
			Err:    err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "更新成功",
		Data:   task,
	}
}
