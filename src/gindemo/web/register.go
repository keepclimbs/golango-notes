package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	req := &registerReq{}
	err := c.ShouldBind(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, req)
	}

}

type registerReq struct {
	Phone string `form:"phone"     binding:"required"`
	Email string `form:"email" binding:"omitempty"`
}
