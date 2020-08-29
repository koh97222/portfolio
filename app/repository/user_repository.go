package repository

import (
	"app/domain"
	"app/infrastructure"
)

// CreateNewUser ユーザ情報を新規登録します。
func CreateNewUser(u domain.User) (domain.User, error) {

	db := infrastructure.DbConnect()
	query := "INSERT INTO users (userID,password,name,Email,Year,Month,Day,Sex) VALUES (?,?,?,?,?,?,?,?)"

	res, err := db.Exec(query, u.UserID, u.Password, u.Name, u.Email, u.Year, u.Month, u.Day, u.Sex)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 新規に発番されたidを記録。
	id, _ := res.LastInsertId()
	u.ID = int(id)
	return u, err
}

// IsDuplicateUserID ユーザIDが既に登録されているかどうかを確認します。
func IsDuplicateUserID(u domain.User) (bool, error) {

	db := infrastructure.DbConnect()
	query := "SELECT count(*) FROM users WHERE userID = ?"

	cnt, err := db.Query(query, u.UserID)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// userID検索結果件数
	var Count domain.Cnt
	for cnt.Next() {
		err := cnt.Scan(&Count.Count)
		if err != nil {
			panic(err.Error())
		}
	}

	// ユーザIDが重複している場合、trueを返します。
	return Count.Count != 0, err
}
