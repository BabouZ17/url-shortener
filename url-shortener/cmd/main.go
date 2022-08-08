package main

import (
	"os"

	"github.com/BabouZ17/url-shortener/pkg/config"
	"github.com/BabouZ17/url-shortener/pkg/controller"
	"github.com/BabouZ17/url-shortener/pkg/db"
	"github.com/BabouZ17/url-shortener/pkg/repository"
	"github.com/BabouZ17/url-shortener/pkg/router"
	"github.com/gin-gonic/gin"
)

func initialize() {
	r := gin.Default()
	config := config.New(os.Getenv("CONFIG_PATH"))
	dataBase := db.New(config)
	urlRepository := repository.NewUrlRepository(dataBase)
	urlController := controller.NewUrlController(urlRepository)
	statusController := controller.NewStatusController(urlRepository)

	router.StatusRoutes(r, statusController)
	router.UrlRoutes(r, urlController)

	r.Run("0.0.0.0:8080")
}

func main() {
	initialize()
}
