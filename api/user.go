package api

import (
	"github.com/gin-gonic/gin"
	"todo_list/serializer"
	"todo_list/service"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBindJSON(&userService); err == nil {
		res := userService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.ResponseError{Msg: err.Error()})
	}
}

func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBindJSON(&userService); err == nil {
		res := userService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, serializer.ResponseError{Msg: err.Error()})
	}
}

func UserGet(c *gin.Context) {
	c.JSON(200, service.GetAll())
}
