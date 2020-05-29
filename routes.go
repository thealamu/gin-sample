package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes() {
	router.GET("/", showIndexPage)
}

func showIndexPage(c *gin.Context) {
	c.HTML(
		// 200 OK
		http.StatusOK,

		// Use the index html template
		"index.html",

		// Pass the data the template uses
		gin.H{
			"title": "Home Page",
		})
}
