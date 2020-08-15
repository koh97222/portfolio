package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// User ユーザ情報の構造体
type User struct {
	ID    int    `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Year  int    `gorm:"column:year"`
	Month int    `gorm:"column:month"`
	Day   int    `gorm:"column:day"`
	Sex   int    `gorm:"column:sex"`
	gorm.Model
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

// registNewUser 疎通確認用メソッド
func registNewUser() User {
	db := gormConnect()
	sample := User{}
	db.Where("id = ?", "1").Find(&sample)
	defer db.Close()
	return sample
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/registUser", func(c echo.Context) error {
		return c.JSON(http.StatusOK, registNewUser())
	})
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8082"))
	// e.Startの中はdocker-cogitmposeのgoコンテナで設定したportsを指定してください。
}
