package routers

import (
	"fmt"
	"gin/middleware"
	"gin/web"

	"github.com/gin-gonic/gin"
)

func initCourse(r *gin.Engine) {
	fmt.Println("InitCourse")
	// http://localhost:8080/v1
	course := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)
	course.POST("/course", web.CreateCourse)
	course.GET("/course", web.GetCourse)
}
