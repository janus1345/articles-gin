package tests

import (
	"articles/handlers"
	"articles/models"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnAuthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/", handlers.ShowIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
		return statusOk && pageOk
	})
}

func TestShowIndexPageAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)

	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	r.GET("/", handlers.ShowIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Home Page</title>") < 0 {
		t.Fail()
	}
}

func TestArticleUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", handlers.GetArticle)
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>article 1</title>") > 0
		return statusOk && pageOk
	})
}

func TestArticleListJson(t *testing.T) {
	r := getRouter(true)
	r.GET("/", handlers.ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []models.Article
		err = json.Unmarshal(p, &articles)
		return err == nil && len(articles) >= 2 && statusOk

	})
}

func TestArticleListXML(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/view/:article_id", handlers.GetArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var article models.Article
		err = xml.Unmarshal(p, &article)
		fmt.Println(err)
		return statusOk && err == nil && article.Id == 1 && article.Title == "article 1"
	})
}



