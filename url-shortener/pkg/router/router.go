package router

import (
	"github.com/BabouZ17/url-shortener/pkg/controller"
	"github.com/gin-gonic/gin"
)

func UrlRoutes(router *gin.Engine, controller *controller.UrlController) *gin.Engine {
	url := router.Group("/")
	{
		url.GET("/urls", controller.ListUrls)
		url.POST("urls", controller.AddUrl)
		url.GET("/urls/:id", controller.GetUrlById)
		url.DELETE("urls/:id", controller.DeleteUrl)
		url.GET("/redirect/:alias", controller.RedirectToTarget)
	}
	return router
}

func StatusRoutes(router *gin.Engine, controller *controller.StatusController) *gin.Engine {
	status := router.Group("/status")
	{
		status.GET("/readiness", controller.Readiness)
		status.GET("/liveness", controller.Liveness)
	}
	return router
}
