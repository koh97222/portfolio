package logic

import (
	"app/domain"
	"app/repository"

	"github.com/labstack/echo"
)

// CreateNewUser 新規にユーザを登録します。
func CreateNewUser(c echo.Context) (res *domain.Response, err error) {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		res = res.Failed("invalid parameter.", nil)
		return res, err
	}
	// userIdの重複がないかどうかを確認します。
	// 重複がない場合は、そのまま登録に進みます。
	isDuplicated, err := repository.IsDuplicateUserID(*user)
	if isDuplicated {
		res = res.Failed("user has not created because userID is already exist.", nil)
		return res, err
	}
	// ユーザ登録メソッド実行
	*user, err = repository.CreateNewUser(*user)
	if err != nil {
		res = res.Failed("user registration is failed.", nil)
		return res, err
	}
	res = res.Success("user has created.", *user)
	return res, err
}
