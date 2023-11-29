// ユーザーの認識を行う関数の作成
package model

import (
	"errors"
	"financialManagementApplication/db"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username, password string) (*TbTAccount, error) {
	var user TbTAccount
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Error during password comparison: %v", err)
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func GetHashedPasswordByUsername(username string) (string, error) {
	log.Printf("GetHashedPasswordByUsername() is called")
	var user TbTAccount
	if result := db.DB.Where("username = ?", username).First(&user); result.Error != nil {
		return "", errors.New("Error retrieving user from database")
	}
	return user.Password, nil
}
