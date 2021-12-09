package handlers

import (
	"articles/models"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func ShowRegistrationPage(c *gin.Context) {
	render(c, gin.H{"title": "Register"}, "register.html")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if _, err := models.RegisterNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, false)
		c.Set("is_logged_in", "true")
		render(c, gin.H{"title": "Successful registration && Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		})
	}
}
