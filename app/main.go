package main

import (
	"app/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.POST("/createNewUser", controller.CreateNewUser) // ユーザー新規登録API
	e.POST("/login", controller.Login)                 // ログインAPI
	e.GET("/trialLogin", controller.TrialLogin)        // お試しログインAPI
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8082"))
}
