package controller

import (
	"app/logic"
	"net/http"

	"github.com/labstack/echo"
)

// CreateNewUser 新規ユーザ登録ハンドラ。
func CreateNewUser(c echo.Context) error {
	res, err := logic.CreateNewUser(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
