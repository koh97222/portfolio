package repository

import (
	"app/domain"
	"app/infrastructure"
)

// GetUserInfo ユーザ情報の取得
func GetUserInfo(userID string) (domain.User, error) {

	var user domain.User

	db := infrastructure.DbConnect()
	query := "SELECT * FROM users WHERE userID = ? "

	res, err := db.Query(query, userID)
	if err != nil {
		panic(err.Error())
	}

	for res.Next() {
		err := res.Scan(&user.ID, &user.UserID, &user.Password, &user.Name, &user.Email, &user.Year, &user.Month, &user.Day, &user.Sex)
		if err != nil {
			panic(err.Error())
		}
	}
	return user, err
}

// VerifyInput ユーザIDとパスワードの照合を行う。
func VerifyInput(u *domain.UserLoginInput) (bool, error) {

	db := infrastructure.DbConnect()
	query := "SELECT count(*) FROM users WHERE userID = ? and password = ?"

	cnt, err := db.Query(query, u.UserID, u.Password)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// userID,password照合結果確認
	var Count domain.Cnt
	for cnt.Next() {
		err := cnt.Scan(&Count.Count)
		if err != nil {
			panic(err.Error())
		}
	}
	return Count.Count != 0, err
}
