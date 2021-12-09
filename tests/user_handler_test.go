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
}

func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
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
