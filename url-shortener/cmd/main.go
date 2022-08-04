package main

import (
	"github.com/gin-gonic/gin"
	"github.com/BabouZ17/url-shortener/pkg/router"
)

func initialize() {
	r := gin.Default()
	router.SetRoutes(r)
	r.Run()
}
  
func main() {
	initialize()
}