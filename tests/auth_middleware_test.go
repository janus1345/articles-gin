package tests

import (
	"articles/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnsureLoginUnauthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), middleware.EnsureLogin(), func(c *gin.Context) {
		t.Fail()
	})
	testMiddlewareRequest(t, r, http.StatusUnauthorized)

}

func TestEnsureLoginInAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), middleware.EnsureLogin(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoginUnAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(false), middleware.EnsureNotLogin(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	testMiddlewareRequest(t, r, http.StatusOK)
}

func TestEnsureNotLoginAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", setLoggedIn(true), middleware.EnsureNotLogin(), func(c *gin.Context) {
		t.Fail()
	})
	testMiddlewareRequest(t, r, http.StatusUnauthorized)
}

func TestSetUserStatusAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", middleware.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		if !exists || !loggedInInterface.(bool) {
			t.Fail()
		}
	})
	w := httptest.NewRecorder()
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)

}

func TestSetUserStatusUnAuthenticated(t *testing.T) {
	r := getRouter(false)
	r.GET("/", middleware.SetUserStatus(), func(c *gin.Context) {
		loggedInInterface, exists := c.Get("is_logged_in")
		fmt.Println(loggedInInterface.(bool), exists)
		if exists && loggedInInterface.(bool) {
			t.Fail()
		}
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

}

func setLoggedIn(b bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_logged_in", b)
	}
}
