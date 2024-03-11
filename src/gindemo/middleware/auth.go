package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userName, _ := c.Get("user_name")
	fmt.Printf("auth check userId:%s, userName: %s", userId, userName)
	c.Next()
}

var token = "123456"

func TokenCheck(c *gin.Context) {
	accessToken := c.Request.Header.Get("access_token")
	if accessToken != token {
		c.JSON(http.StatusOK, gin.H{
			"message": "token not equal",
		})
		c.AbortWithError(http.StatusInternalServerError, errors.New("token 检查失败")) // 上下文信息 个人理解是这个
	}
	c.Set("user_name", "nick")
	c.Set("user_id", "10001")
	c.Next()
}
