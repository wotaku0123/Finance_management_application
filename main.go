package main

import (
	"financialManagementApplication/db"
	// "sokavodHservice/routes"
	"financialManagementApplication/controllers"

	"github.com/kataras/iris/v12"
)

func main() {

	db.Init()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Post("/login", controllers.Login)

	app.Run(iris.Addr(":8080"))
}
