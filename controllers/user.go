package controllers

import (
	"log"

	"financialManagementApplication/libs"
	"financialManagementApplication/model"

	"github.com/kataras/iris/v12"
)

func CreateNewUserAccount(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)

	user := new(model.TbTAccount)
	if err := ctx.ReadJSON(user); err != nil {
		_ = ctx.JSON(libs.ApiResource(2, nil, err.Error()))
		return
	}

	if err := user.CreateUser(); err != nil {
		errMsg := "ユーザの登録に失敗しました。時間をおいてもう一度試してください。"
		log.Println(errMsg, user)
		_ = ctx.JSON(libs.ApiResource(3, nil, err.Error()))
		return
	}
	_ = ctx.JSON(libs.ApiResource(1, nil, "成功"))
}

// func Login(ctx iris.Context) {
// 	ctx.StatusCode(iris.StatusOK)

// 	aul := new(model.TbTAccount)
// 	if err := ctx.ReadJSON(aul); err != nil {
// 		_ = ctx.JSON(libs.ApiResource(2, nil, err.Error()))
// 		return
// 	}
// 	user, err := model.GetUser(int(aul.ID))
// 	if err != nil {
// 		_ = ctx.JSON(libs.ApiResource(3, nil, err.Error()))
// 		return
// 	}
// 	user.CheckLogin(ctx, aul.Password)
// 	_ = ctx.JSON(libs.ApiResource(1, user, "成功"))
// }

func VerifyTest(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK)
	_ = ctx.JSON(libs.ApiResource(1, nil, "成功"))
}
