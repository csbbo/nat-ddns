package middleware

import (
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//session := sessions.Default(c)
		//filter := c.FullPath() == "/login" || c.FullPath() == "/regist" || c.FullPath() == "/welcome"
		//if name, ok:=session.Get("user").(string); (!ok || name== "") && !filter{
		//	common.ResponseError(c, "用户未登录")
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
