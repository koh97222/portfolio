package domain

// User ユーザ情報のモデル構造体
type User struct {
	ID       int
	UserID   string
	Password string
	Name     string
	Email    string
	Year     int
	Month    int
	Day      int
	Sex      string
}

// Users 複数のユーザ情報のモデル構造体
type Users []User

// UserLoginInput ユーザログイン入力のモデル構造体
type UserLoginInput struct {
	UserID   string
	Password string
}
