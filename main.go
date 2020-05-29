package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Use the default router
	router = gin.Default()

	// Process the templates
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			// 200 OK
			http.StatusOK,

			// Use the index html template
			"index.html",

			// Pass the data the template uses
			gin.H{
				"title": "Home Page",
			})
	})

	router.Run()
}
