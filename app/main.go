package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/connect", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-cogitmposeのgoコンテナで設定したportsを指定してください。
}
