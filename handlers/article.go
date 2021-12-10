package handlers

import (
	"articles/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func render(c *gin.Context, data gin.H, templateName string) {
	//loogedInInterface, _ := c.Get("is_logged_in")
	//data["is_logged_in"] = loogedInInterface.(bool)
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}

}

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	//c.HTML(
	//	http.StatusOK,
	//	"index.html",
	//	gin.H{
	//		"title": "Home Page",
	//		"payload": articles,
	//	},
	//	)
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	},
		"index.html",
	)
}

func GetArticle(c *gin.Context) {
	articleId := c.Param("article_id")
	if id, err := strconv.Atoi(articleId); err == nil {
		if article, err := models.GetArticleById(id); err == nil {
			//c.HTML(
			//	http.StatusOK,
			//	"article.html",
			//	gin.H{
			//		"title": article.Title,
			//		"payload": article,
			//	},
			//)
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			},
				"article.html",
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}

func ShowArticleCreationPage(c *gin.Context) {
	render(c, gin.H{"title": "Create New Article"}, "create-article.html")

}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := models.CreateNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a,
		}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
