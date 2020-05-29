package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initRoutes() {
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

}

func getArticle(c *gin.Context) {
	// Get article id
	if article_id, err := strconv.Atoi(c.Param("article_id")); err == nil {

		if article, err := getArticleById(article_id); err == nil {

			c.HTML(
				http.StatusOK,

				"article.html",

				gin.H{
					"title":   article.Title,
					"payload": article,
				})
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func getArticleById(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	c.HTML(
		// 200 OK
		http.StatusOK,

		// Use the index html template
		"index.html",

		// Pass the data the template uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		})
}
