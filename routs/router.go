package routs

import (
	"github.com/gin-gonic/gin"
	"todo_list/api"
	"todo_list/middleware"
)

func CreateRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
	}
	authed := v1.Group("/")
	authed.Use(middleware.JWT())
	{
		authed.GET("/user", api.UserGet)
	}
	return router
}
