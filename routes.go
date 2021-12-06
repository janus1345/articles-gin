package main

import "articles/handlers"

func InitializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
}