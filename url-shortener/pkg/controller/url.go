package controller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/BabouZ17/url-shortener/pkg/model"
	"github.com/BabouZ17/url-shortener/pkg/repository"
	"github.com/BabouZ17/url-shortener/pkg/service"
	"github.com/gin-gonic/gin"
)

type UrlController struct {
	repository *repository.UrlRepository
}

func NewUrlController(repository *repository.UrlRepository) *UrlController {
	return &UrlController{repository}
}

func (uc UrlController) AddUrl(c *gin.Context) {
	var newUrl model.UrlCreationDTO
	if err := c.ShouldBindJSON(&newUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if _, err := url.Parse(newUrl.Target); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	url := model.NewUrl(service.RandomString(6), newUrl.Target)
	url, err := uc.repository.Add(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, url)
}

func (uc UrlController) ListUrls(c *gin.Context) {
	urls, err := uc.repository.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, urls)
}

func (uc UrlController) GetUrlById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid id received"})
		return
	}

	url, err := uc.repository.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, url)
}

func (uc UrlController) DeleteUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid id received"})
		return
	}

	if err := uc.repository.Delete(id); err != nil {
		switch err.(type) {
		case *repository.RepositoryError:
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (uc UrlController) RedirectToTarget(c *gin.Context) {
	alias := c.Param("alias")

	url, err := uc.repository.GetByAlias(alias)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, url.Target)
}
