package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	fmt.Println("InitRouters")
	initApi(r)
	initCourse(r)
}
