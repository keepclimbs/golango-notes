package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "yes",
	})
}

func EditCourse(c *gin.Context) {

}

func DeleteCourse(c *gin.Context) {

}

func GetCourse(c *gin.Context) {

}
