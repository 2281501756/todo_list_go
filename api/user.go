package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo_list/serializer"
	"todo_list/service"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Register()
		c.JSON(200, res)
	} else {
		fmt.Println(err)
		c.JSON(400, serializer.ResponseError{Msg: err.Error()})
	}
}
