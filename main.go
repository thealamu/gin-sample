package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Use the default router
	router = gin.Default()

	// Process the templates
	router.LoadHTMLGlob("templates/*")

	// Handle routes
	initRoutes()

	router.Run()
}
