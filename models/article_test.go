package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, a := range alist {
		if a.ID != articleList[i].ID ||
			a.Title != articleList[i].Title ||
			a.Content != articleList[i].Content {
			t.Fail()
			break
		}
	}

}
