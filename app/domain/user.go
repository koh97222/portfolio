package domain

// User ユーザ情報のモデル構造体
type User struct {
	ID       int    `gorm:"column:id"`
	UserID   string `gorm:"column:userId"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Year     int    `gorm:"column:year"`
	Month    int    `gorm:"column:month"`
	Day      int    `gorm:"column:day"`
	Sex      string `gorm:"column:sex"`
}

// Users 複数のユーザ情報のモデル構造体
type Users []User
