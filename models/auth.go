package models

import (
	"golang.org/x/crypto/bcrypt"
)

/*Auth structure manages auth model */
type Auth struct{}

/*Sign return user login */
func (Model *Auth) Sign(User *User) bool {
	password := []byte(User.Password)
	exists := User.ExistsEmail()
	if !exists {
		return false
	}
	passwordEncrypt := []byte(User.Password)
	err := bcrypt.CompareHashAndPassword(passwordEncrypt, password)
	if err != nil {
		return false
	}
	return true
}
