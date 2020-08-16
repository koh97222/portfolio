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

// User ユーザ情報の構造体
type User struct {
	UserID   string `gorm:"column:userId"`
	Password string `gorm:"column:password"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Year     int    `gorm:"column:year"`
	Month    int    `gorm:"column:month"`
	Day      int    `gorm:"column:day"`
	Sex      string `gorm:"column:sex"`
}

// Response APIのレスポンス結果に対応する汎用的な構造体
type Response struct {
	ResultCode    string `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
	User
}

// Response構造体の初期化
// 基本的に処理失敗時のみ、ResultCodeを変更するようにしたい。
func initResponse() *Response {
	r := new(Response)
	r.ResultCode = SuccessCode
	r.ResultMessage = ""
	return r
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
func registNewUser(u User) {
	db := gormConnect()
	db.Create(u)
	defer db.Close()
}

// registNewUserHandler 新規ユーザ登録ハンドラ。
func registNewUserHandler(c echo.Context) error {
	r := initResponse()
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	registNewUser(*user)
	r.ResultMessage = "registerd"
	return c.JSON(http.StatusOK, r)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/registUser", func(c echo.Context) error {
		return c.JSON(http.StatusOK, registNewUserHandler(c))
	})
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-cogitmposeのgoコンテナで設定したportsを指定してください。
}
