package models

import (
	"golang.org/x/crypto/bcrypt"
)

// Auth struct
type Auth struct{}

// Authentication validate access authentication
func (_Auth *Auth) Authentication(_User *User) bool {
	password := []byte(_User.Password)

	exists := _User.CheckExists()
	if !exists {
		return false
	}

	passwordEncrypt := []byte(_User.Password)
	err := bcrypt.CompareHashAndPassword(passwordEncrypt, password)
	if err != nil {
		return false
	}

	return true
}
