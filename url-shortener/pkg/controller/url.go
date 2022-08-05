package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/BabouZ17/url-shortener/pkg/model"
	"github.com/BabouZ17/url-shortener/pkg/repository"
)

type UrlController struct {
	repository *repository.UrlRepository
}

func New(repository *repository.UrlRepository) *UrlController {
	return &UrlController{repository}
}

func (uc UrlController) ListUrls(c *gin.Context) {
	var urls []model.Url
	urls = uc.repository.List()
	c.JSON(http.StatusOK, urls)
}

func (uc UrlController) GetUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid id received"})
		return
	}

	url, err := uc.repository.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)
}