package main

import (
	"articles/handlers"
)

func InitializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/register", handlers.ShowRegistrationPage)
		userRoutes.POST("/register", handlers.Register)

		userRoutes.GET("/login", handlers.ShowLoginPage)
		userRoutes.POST("/login", handlers.PerformLogin)
		userRoutes.GET("/logout", handlers.Logout)
	}
}
