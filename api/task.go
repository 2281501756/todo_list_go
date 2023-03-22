package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo_list/serializer"
	"todo_list/service"
)

func TaskGetALL(c *gin.Context) {
	res := service.GetALlTask()
	c.JSON(200, res)
}

func TaskCreate(c *gin.Context) {
	var taskCreateService service.TaskCreateService
	if err := c.ShouldBindJSON(&taskCreateService); err == nil {
		id, has := c.Get("id")
		if !has {
			c.JSON(400, serializer.ResponseError{Msg: "未登录"})
		}
		res := taskCreateService.Create(id.(uint))
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.ResponseError{Msg: err.Error()})
	}
}
func TaskUpdate(c *gin.Context) {
	var taskUpdateService service.TaskUpdateService
	if err := c.ShouldBindJSON(&taskUpdateService); err == nil {
		res := taskUpdateService.Update()
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.ResponseError{Msg: err.Error()})
	}
}
func TaskDelete(c *gin.Context) {
	type data struct {
		Id uint `json:"id" form:"id" binding:"required"`
	}
	var d data
	if err := c.ShouldBindJSON(&d); err == nil {
		fmt.Println(d.Id)
		c.JSON(200, service.DeleteTask(d.Id))
	} else {
		c.JSON(400, serializer.ResponseError{Msg: err.Error()})
	}
}
