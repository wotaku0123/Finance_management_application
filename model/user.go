package model

import (
	"financialManagementApplication/db"
	"log"

	"github.com/jameskeane/bcrypt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

// / CheckLogin check login user
func (u *TbTAccount) CheckLogin(ctx iris.Context, password string) (int64, *jwt.TokenPair, string) {
	if u.ID == 0 {
		return 2, nil, "ユーザが存在しません"
	} else {
		if bcrypt.Match(password, u.Password) {
			tokenPair := u.GenerateTokenPair(ctx)
			return 1, tokenPair, "ログインに成功しました"
		} else {
			return 3, nil, "ログイン情報が正しくありません"
		}
	}
}

func (u *TbTAccount) CreateUser() error {
	log.Println("CreateUser function called")
	var err error
	u.Password, err = HashPassword(u.Password)
	if err != nil {
		return err
	}
	if err := db.DB.Create(u).Error; err != nil {
		return err
	}
	log.Printf("Hashed password for %s: %s", u.Name, u.Password)
	return nil
}
func HashPassword(pwd string) (string, error) {
	salt, err := bcrypt.Salt(10)
	if err != nil {
		return "", err
	}
	hash, err := bcrypt.Hash(pwd, salt)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func GetUser(id int) (*TbTAccount, error) {
	var user TbTAccount
	err := db.DB.First(&user, id).Error
	return &user, err
}
