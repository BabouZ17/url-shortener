package router

import (
	"github.com/gin-gonic/gin"
	"github.com/BabouZ17/url-shortener/pkg/controller"
)

func SetRoutes(router *gin.Engine) *gin.Engine {
	api := router.Group("/")
	{
		api.GET("/urls", controller.ListUrls)
		api.GET("/urls/:id", controller.GetUrl)
		api.POST("urls", controller.AddUrl)
		api.DELETE("urls/:id", controller.DeleteUrl)
		api.DELETE("urls", controller.DeleteUrls)
	}
	status := router.Group("/status")
	{
		status.GET("/health", controller.Alive)
	}
	return router
}
