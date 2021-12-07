package models

import (
	"errors"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var UserList = []User{
	{"user1", "pass1"},
	{"user2", "pass2"},
	{"user3", "pass3"},
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if isUsernameAviailable(username) {
		return nil, errors.New("The username isn't available")
	}
	u := User{Username: username, Password: password}
	UserList = append(UserList, u)
	return &u, nil
}

func isUsernameAviailable(username string) bool {
	for _, v := range UserList {
		if v.Username == username {
			return true
		}
	}
	return false
}
