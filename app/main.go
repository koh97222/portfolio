package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// 定数
const (
	SuccessCode = "00" // 処理成功時
	FailedCode  = "80" // 処理失敗時
)

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

// UserResponse ユーザ情報のレスポンス構造体
type UserResponse struct {
	Response Response `json:"response"`
	User     *User    `json:"user"`
}

// Response APIの処理結果に対応する汎用的な構造体
type Response struct {
	ResultCode    string `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

// GormConnect Mysqlへ接続し、dbインスタンスを返却します。
func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "portfolio_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// registNewUser ユーザ情報を新規登録します。
func registNewUser(u User) *User {
	db := gormConnect()
	db.Create(&u)
	defer db.Close()
	return &u
}

// isDuplicateUserID ユーザIDが既に登録されているかどうかを確認します。
func isDuplicateUserID(u User) bool {
	db := gormConnect()
	user := User{}
	isNotFound := db.Where("userID = ?", u.UserID).Find(&user).RecordNotFound()
	defer db.Close()
	return !isNotFound
}

// registNewUserHandler 新規ユーザ登録ハンドラ。
func registNewUserHandler(c echo.Context) error {
	user := new(User)
	res := new(UserResponse)
	if err := c.Bind(user); err != nil {
		return err
	}
	// userIdの重複がないかどうかを確認します。
	// 重複がない場合は、そのまま登録に進みます。
	if isDuplicateUserID(*user) {
		res.User = nil
		res.Response.ResultCode = FailedCode
		res.Response.ResultMessage = "user has not created because userID is already exist."
		return c.JSON(http.StatusOK, user)
	}
	// ユーザ登録メソッド実行
	user = registNewUser(*user)
	res.User = user
	res.Response.ResultCode = SuccessCode
	res.Response.ResultMessage = "user has created."

	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()
	// ユーザー新規登録API
	e.POST("/registUser", registNewUserHandler)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8082"))
}
