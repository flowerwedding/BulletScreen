package util

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CompareTag() gin.HandlerFunc {
	return func(c *gin.Context) {
        t := sessions.Default(c)
        tag := t.Get("tag")
        fmt.Println(tag)
		if tag == "用户已禁止"{
			c.Abort()
			c.JSON(http.StatusUnauthorized, "Token is not valid")
		}else if tag == "账号未验证"{
			c.Abort()
			c.JSON(http.StatusUnauthorized, "Token is not valid")
		}
		c.Next()
	}
}
