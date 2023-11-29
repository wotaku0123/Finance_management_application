package model

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"financialManagementApplication/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

const (
	accessTokenMaxAge  = 7 * 24 * time.Hour
	refreshTokenMaxAge = 14 * 24 * time.Hour
	leewayAge          = 8 * 24 * time.Hour
)

var (
	privateKey, publicKey = jwt.MustLoadRSA("rsa_private_key.pem", "rsa_public_key.pem")

	signer   = jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)
	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)

type UserClaims struct {
	ID string `json:"user_id"`
	// Do: `json:"username,required"` to have this field required
	// or see the Validate method below instead.
	Name string `json:"username"`
	exp  int64
	iat  int64
}

// generate token to user and store token
func (u *TbTAccount) GenerateTokenPair(ctx iris.Context) *jwt.TokenPair {

	var tempID string

	tempID = strconv.Itoa(int(u.ID))
	refreshClaims := jwt.Claims{Subject: tempID}

	accessClaims := UserClaims{
		ID:   tempID,
		Name: u.Name,
		exp:  time.Now().UTC().Add(time.Hour * time.Duration(1)).Unix(),
		iat:  time.Now().UTC().Unix(),
	}
	tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		ctx.Application().Logger().Errorf("token pair: %v", err)
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return nil
	}
	u.AccessToken, _ = strconv.Unquote(string(tokenPair.AccessToken))
	u.RefreshToken, _ = strconv.Unquote(string(tokenPair.RefreshToken))

	fmt.Println(u.AccessToken)
	ups := map[string]interface{}{"access_token": u.AccessToken, "refresh_token": u.RefreshToken}
	if err := db.DB.Model(&TbTAccount{}).Where("id = ?", u.ID).Updates(ups).Error; err != nil {
		log.Println(fmt.Sprintf("Token保存失敗：%+v\n", u.ID))
		log.Println(fmt.Sprintf("Token保存失敗：%+v\n", err))
	}
	return &tokenPair

}
