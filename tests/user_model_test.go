package tests

import (
	"articles/models"
	"testing"
)

func TestValidUserRegistration(t *testing.T) {
	saveLists()
	if r, err := models.RegisterNewUser("test", "test"); err == nil {
		if r.Username != "test" || r.Password != "test" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
	restoreList()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()
	if u, err := models.RegisterNewUser("user1", "pass1"); err == nil || u != nil {
		t.Fail()
	}
	if u, err := models.RegisterNewUser("newuser", ""); err == nil || u != nil {
		t.Fail()
	}
	restoreList()
}
