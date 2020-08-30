package logic

import (
	"app/domain"
	"app/repository"

	"github.com/labstack/echo"
)

// Login ログインロジック
// TODO:今は簡易的な実装なので、後々なれてきたらちゃんと実装する。
func Login(c echo.Context) (res *domain.Response, err error) {
	input := new(domain.UserLoginInput)
	if err := c.Bind(input); err != nil {
		res = res.Failed("invalid parameter.", nil)
		return res, err
	}

	// 入力されたUserID,Passwordが登録済みのユーザ情報であるかを確認します。
	verifyok, err := repository.VerifyInput(input)
	if err != nil {
		res = res.Failed("an error occur when user input verifying.", nil)
		return res, err
	}
	// ログイン失敗時は、それを伝えます。
	if !verifyok {
		res = res.Failed("user verification failed.", nil)
		return res, err
	}
	// 登録済みの場合は、ユーザ情報を返却します。
	// ユーザ情報取得
	user, err := repository.GetUserInfo(input.UserID)
	if err != nil {
		res = res.Failed("an error occur when user info acquisition.", nil)
		return res, err
	}

	res = res.Success("login success.", user)
	return res, err
}

// TrialLogin お試しログインロジック
func TrialLogin(c echo.Context) (res *domain.Response, err error) {
	// お試し用ユーザの情報を取得して返却します。
	user, err := repository.GetUserInfo("trial_user")
	if err != nil {
		res = res.Failed("an error occur when trial user info aquisition.", nil)
	}

	res = res.Success("trial login success.", user)
	return res, err
}
