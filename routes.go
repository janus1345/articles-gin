package main

import (
	"articles/handlers"
	"articles/middleware"
)

func InitializeRoutes() {
	router.Use(middleware.SetUserStatus())
	router.GET("/", handlers.ShowIndexPage)

	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/register", middleware.EnsureNotLogin(), handlers.ShowRegistrationPage)
		userRoutes.POST("/register", middleware.EnsureNotLogin(), handlers.Register)

		userRoutes.GET("/login", middleware.EnsureNotLogin(), handlers.ShowLoginPage)
		userRoutes.POST("/login", middleware.EnsureNotLogin(), handlers.PerformLogin)
		userRoutes.GET("/logout", middleware.EnsureLogin(), handlers.Logout)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)
		articleRoutes.GET("/create", middleware.EnsureLogin(), handlers.ShowArticleCreationPage)
		articleRoutes.POST("/create", middleware.EnsureLogin(), handlers.CreateArticle)
	}
}
