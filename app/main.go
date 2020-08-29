package main

import (
	"app/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// ユーザー新規登録API
	e.POST("/registUser", controller.CreateNewUser)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8082"))
}
