package routs

import (
	"github.com/gin-gonic/gin"
	"todo_list/api"
)

func CreateRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
	}
	return router
}
