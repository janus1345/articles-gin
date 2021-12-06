package tests

import (
	"articles/models"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	alist := models.GetAllArticles()
	articleList := models.ArticleList
	if len(alist) != len(articleList) {
		t.Fail()
	}
	for i, v := range alist {
		if v.Id != articleList[i].Id || v.Content != articleList[i].Content || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
