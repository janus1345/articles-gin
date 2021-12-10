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

func TestGetArticleByID(t *testing.T) {
	a, err := models.GetArticleById(1)
	if err != nil || a.Id != 1 || a.Title != "article 1" || a.Content != "article 1 body" {
		t.Fail()
	}
}

func TestCreateNewArticle(t *testing.T) {
	saveLists()
	originalLength := len(models.GetAllArticles())
	a, err := models.CreateNewArticle("new test title", "new test content")
	allArticles := models.GetAllArticles()
	newLength := len(allArticles)
	if err != nil || newLength != originalLength+1 || a.Title != "new test title" || a.Content != "new test content" {
		t.Fail()
	}
	restoreList()
}
