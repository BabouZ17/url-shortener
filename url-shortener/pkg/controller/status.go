package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Alive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "alive",
	})
}
