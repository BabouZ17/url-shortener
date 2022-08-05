package main

import (
	"os"
	"github.com/BabouZ17/url-shortener/pkg/controller"
	"github.com/BabouZ17/url-shortener/pkg/db"
	"github.com/BabouZ17/url-shortener/pkg/repository"
	"github.com/BabouZ17/url-shortener/pkg/router"
	"github.com/BabouZ17/url-shortener/pkg/config"
	"github.com/gin-gonic/gin"
)

func initialize() {
	r := gin.Default()
	config := config.New(os.Getenv("CONFIG_PATH"))
	dataBase := db.New(config)
	urlRepository := repository.New(dataBase)
	urlController := controller.New(urlRepository)

	router.StatusRoutes(r)
	router.UrlRoutes(r, urlController)

	r.Run()
}

func main() {
	initialize()
}
