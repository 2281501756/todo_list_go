package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"todo_list/package/utils"
	"todo_list/serializer"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if claim.ExpiresAt < time.Now().Unix() {
				code = 401
			}
			if claim == nil {
				c.JSON(200, serializer.Response{Msg: "token信息错误", Status: code})
				c.Abort()
				return
			}
			c.Set("id", claim.Id)
		}
		if code != 200 {
			c.JSON(200, serializer.Response{Msg: "token信息错误", Status: code})
			c.Abort()
			return
		}
		c.Next()
	}
}
