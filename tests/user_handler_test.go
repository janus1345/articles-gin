package tests

import (
	"articles/handlers"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/user/register", handlers.ShowRegistrationPage)
	req, _ := http.NewRequest("GET", "/user/register", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Register</title>") > 0
		return statusOk && pageOk
	})
}

func TestRegisterUnauthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)
	r.POST("/user/register", handlers.Register)
	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/user/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	fmt.Println(string(p))
	if err != nil || strings.Index(string(p), "<title>Successful registration &amp;&amp; Login</title>") < 0 {
		t.Fail()
	}
	restoreList()
}

func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()
	r := getRouter(true)
	r.POST("/user/register", handlers.Register)
	registrationPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/user/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))
	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
	restoreList()
}

func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")
	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")
	return params.Encode()
}

func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/user/login", handlers.ShowLoginPage)

	req, _ := http.NewRequest("GET", "/user/login", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil || strings.Index(string(p), "<title>Login</title>") > 0
		return statusOk && pageOk
	})
}

func TestLoginUnauthenticated(t *testing.T) {
	saveLists()
	w := httptest.NewRecorder()

	r := getRouter(true)
	r.POST("/user/login", handlers.PerformLogin)

	loginPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))
	r.ServeHTTP(w, req)
	fmt.Println(w.Code)
	if w.Code != http.StatusOK {
		t.Fail()
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful login</title>") < 0 {
		t.Fail()
	}
	restoreList()
}

func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	saveLists()

	w := httptest.NewRecorder()

	r := getRouter(true)
	r.POST("/user/login", handlers.PerformLogin)
	loginPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		fmt.Println("***************", w.Code)
		t.Fail()
	}
	restoreList()
}

func TestLogout(t *testing.T) {
	r := getRouter(true)
	r.GET("/user/logout", handlers.Logout)
	req, _ := http.NewRequest("GET", "/user/logout", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statOk := w.Code == http.StatusTemporaryRedirect

		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "Temporary Redirect") > 0
		return statOk && pageOk
	})
}
