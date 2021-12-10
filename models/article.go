package models

import "errors"

type Article struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
}

var ArticleList = []Article{
	{1, "article 1 body", "article 1"},
	{2, "article 2 body", "article 2"},
}

func GetAllArticles() []Article {
	return ArticleList
}

func GetArticleById(id int) (*Article, error) {
	for _, v := range ArticleList {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("Article not found")
}
