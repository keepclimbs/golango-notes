package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	req := loginForm{}
	err := c.ShouldBind(&req) // ShouldBind 好点 如果模型不匹配会返回err
	if err != nil {
		fmt.Println("+++")
		fmt.Println(err)
		fmt.Println("---")

		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, req)
	}
}

type loginForm struct {
	User     string `form:"user" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
