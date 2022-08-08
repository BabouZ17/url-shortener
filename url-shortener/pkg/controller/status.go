package controller

import (
	"net/http"

	"github.com/BabouZ17/url-shortener/pkg/repository"
	"github.com/gin-gonic/gin"
)

type StatusController struct {
	repository *repository.UrlRepository
}

func NewStatusController(repository *repository.UrlRepository) *StatusController {
	return &StatusController{repository}
}

func (sc StatusController) Readiness(c *gin.Context) {
	_, err := sc.repository.Count()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "alive",
	})
}

func (sc StatusController) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "alive"})
}
