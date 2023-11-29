package routes

// クライアント側でログインを行うために、routes.goに以下のルートハンドラーを追加します。

// routes.go
// router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

import (
	"financialManagementApplication/controllers"
	"financialManagementApplication/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

var (
	privateKey, publicKey = jwt.MustLoadRSA("rsa_private_key.pem", "rsa_public_key.pem")

	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)

func App(api *iris.Application) {

	app := api.Party("/").AllowMethods(iris.MethodOptions)
	{
		vue := app.Party("/vue")
		{
			vue.Post("/login", controllers.Login)
			vue.Post("/UserCreate", controllers.CreateNewUserAccount)
			verifyMiddleware := verifier.Verify(func() interface{} {
				return new(model.UserClaims)
			})
			vue.Use(verifyMiddleware)
			vue.Post("/test", controllers.VerifyTest)
		}
	}
}
