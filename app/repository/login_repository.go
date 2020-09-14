package repository

import (
	"app/domain"
	"app/infrastructure"
	"database/sql"
)

// GetUserInfo ユーザ情報の取得
func GetUserInfo(userID string) (domain.User, error) {

	db := infrastructure.DbConnect()

	var user domain.User

	err := infrastructure.Transact(db, func(tx *sql.Tx) error {
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
		return err
	})
	return user, err
}

// VerifyInput ユーザIDとパスワードの照合を行う。
func VerifyInput(u *domain.UserLoginInput) (bool, error) {

	db := infrastructure.DbConnect()

	// userID,password照合結果確認
	var Count domain.Cnt

	err := infrastructure.Transact(db, func(tx *sql.Tx) error {
		query := "SELECT count(*) FROM users WHERE userID = ? and password = ?"
		cnt, err := db.Query(query, u.UserID, u.Password)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		for cnt.Next() {
			err := cnt.Scan(&Count.Count)
			if err != nil {
				panic(err.Error())
			}
		}
		return err
	})

	return Count.Count != 0, err
}
