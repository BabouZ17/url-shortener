package router

import (
	"github.com/gin-gonic/gin"
	"github.com/BabouZ17/url-shortener/pkg/controller"
)

func UrlRoutes(router *gin.Engine, controller *controller.UrlController) *gin.Engine {
	url := router.Group("/")
	{
		url.GET("/urls", controller.ListUrls)
		url.GET("/urls/:id", controller.GetUrl)
		//url.POST("urls", controller.AddUrl)
		//url.DELETE("urls/:id", controller.DeleteUrl)
		//url.DELETE("urls", controller.DeleteUrls)
	}
	return router
}

func StatusRoutes(router *gin.Engine) *gin.Engine {
	status := router.Group("/status")
	{
		status.GET("/health", controller.Alive)
	}
	return router
}
