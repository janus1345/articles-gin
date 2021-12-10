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
	"net/url"
	"strconv"
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

func TestArticleCreationPage(t *testing.T) {
	r := getRouter(true)
	r.GET("/article/create", handlers.ShowArticleCreationPage)
	req, _ := http.NewRequest("GET", "/article/create", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Create New Article</title>") > 0
		return statusOk && pageOk
	})
}

func TestCreateArticle(t *testing.T) {
	saveLists()
	r := getRouter(true)
	r.POST("/article/create", handlers.CreateArticle)
	createArticlePayload := getCreateArticlePayLoad()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(createArticlePayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(createArticlePayload)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "Submission Successful") < 0 {
		t.Fail()
	}
	restoreList()
}

func getCreateArticlePayLoad() string {
	parmas := url.Values{}
	parmas.Add("content", "hahhahah")
	parmas.Add("title", "hhah title")
	return parmas.Encode()
}
