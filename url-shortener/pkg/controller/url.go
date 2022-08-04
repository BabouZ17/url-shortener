package controller

import (
	"net/http"
	"net/url"
	"github.com/gin-gonic/gin"
	"github.com/BabouZ17/url-shortener/pkg/model"
)

var urls = []model.Url{
	model.Url{
		Id: "1",
		Alias: "https://1234.com",
		Target: "https://www.google.com",
	},
}

func GetUrl(c *gin.Context) {
	id := c.Param("id")
	for i, url := range(urls) {
		if url.Id == id {
			c.JSON(http.StatusOK, urls[i])
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Url not found"})
}

func AddUrl(c *gin.Context) {
	var newUrl model.Url
	
	if err := c.ShouldBindJSON(&newUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data received"})
		return
	}

	_, err := url.ParseRequestURI(newUrl.Target)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid target for url"})
		return
	}

	urls = append(urls, newUrl)
	c.JSON(http.StatusCreated, newUrl)
}

func ListUrls(c *gin.Context) {
	c.JSON(http.StatusOK, urls)
}

func DeleteUrl(c *gin.Context) {
	id := c.Param("id")

	for i, url := range(urls) {
		if url.Id == id {
			var newUrls = make([]model.Url, 0)
			var urlDeleted = url

			newUrls = append(newUrls, urls[:i]...)
			newUrls = append(newUrls, urls[i+1:]...)
			urls = newUrls
			
			c.JSON(http.StatusOK, urlDeleted)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Url not found"})
}

func DeleteUrls(c *gin.Context) {
	urls = []model.Url{}
	c.JSON(http.StatusOK, urls)
}