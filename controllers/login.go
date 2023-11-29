package controllers

import (
	"log"
	// "golang.org/x/crypto/bcrypt"
	"financialManagementApplication/model"

	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	//リクエストボディからユーザー名とパスワードを取得
	var request struct {
		Username string `json:"username"` // ダブルクォートの代わりにバッククォートを使用
		Password string `json:"password"` // バッククォートを使用
	}

	log.Println("Login function called")
	allUsers, err := model.GetAllUsers()
	if err != nil {
		log.Println("Error retrieving all users from database:", err)
	} else {
		log.Println("All users from database:")
		for _, user := range allUsers {
			log.Printf("ID: %d, Name: %s, Username: %s, Password: %s\n", user.ID, user.Name, user.Username, user.Password)
		}
	}
	if err := ctx.ReadJSON(&request); err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		return
	}

	userInfo, err := model.GetUserInfo(request.Username)
	if err != nil {
		log.Printf("Error retreiving user info from database: %s\n", err)
		return
	}
	log.Printf("User Info from Database: %+v\n", userInfo)
	log.Printf("Username: %s, Password: %s", request.Username, request.Password)

	user, err := model.Authenticate(request.Username, request.Password)
	if err != nil {
		//この部分のエラーがpostmanで発生している
		log.Println("Error during authentication:", err)
		ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{"error": "Invalid username or password"})
		return
	}

	//データベースからユーザ名に関するハッシュ化されたパスワードを取得しログに出力
	hashedPasswordInDB, err := model.GetHashedPasswordByUsername(request.Username)
	if err != nil {
		log.Printf("Error retrieving hashed password from databse: %s\n", err)
		return
	}

	log.Printf("Hashed Password from Database: %s\n", hashedPasswordInDB)

	tokenPair := user.GenerateTokenPair(ctx)
	if tokenPair == nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(iris.Map{
		"status":        "success",
		"message":       "Logged in successfully",
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}
