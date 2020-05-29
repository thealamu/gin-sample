package main

type article struct {
	ID      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

// Store articles in memory for this sample
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all articles
func getAllArticles() []article {
	return articleList
}
