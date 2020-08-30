package controller

import (
	"app/logic"
	"net/http"

	"github.com/labstack/echo"
)

// Login 登録ユーザログイン
func Login(c echo.Context) error {
	res, err := logic.Login(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, res)
}

// TrialLogin お試しログイン
func TrialLogin(c echo.Context) error {
	res, err := logic.TrialLogin(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
