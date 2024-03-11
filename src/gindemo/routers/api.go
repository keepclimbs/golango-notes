package routers

import (
	"fmt"
	"gin/web"

	"github.com/gin-gonic/gin"
)

func initApi(r *gin.Engine) {
	fmt.Println("InitApi")
	// http://localhost:8080/api
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
	v1.POST("/register1", func(c *gin.Context) {

	})
}
