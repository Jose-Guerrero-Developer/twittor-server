package models

import (
	"golang.org/x/crypto/bcrypt"
)

/*Auth structure manages auth model */
type Auth struct{}

/*Sign return user login */
func (_Auth *Auth) Sign(UserModel *User) bool {
	password := []byte(UserModel.Password)
	exists := UserModel.Exists()
	if !exists {
		return false
	}
	passwordEncrypt := []byte(UserModel.Password)
	err := bcrypt.CompareHashAndPassword(passwordEncrypt, password)
	if err != nil {
		return false
	}
	return true
}
