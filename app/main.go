package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	Result   sql.Result `json:"result"`
	Response Response   `json:"response"`
	User     *User      `json:"user"`
}

// Response APIの処理結果に対応する汎用的な構造体
type Response struct {
	ResultCode    string `json:"resultCode"`
	ResultMessage string `json:"resultMessage"`
}

// Cnt 件数カウント結果
type Cnt struct {
	Count int
}

// GormConnect Mysqlへ接続し、dbインスタンスを返却します。
func dbConnect() *sql.DB {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "portfolio_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := sql.Open(DBMS, CONNECT)
	// db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// registNewUser ユーザ情報を新規登録します。
func registNewUser(u User) sql.Result {

	db := dbConnect()
	query := "INSERT INTO users (userID,password,name,Email,Year,Month,Day,Sex) VALUES (?,?,?,?,?,?,?,?)"
	result, err := db.Exec(query, u.UserID, u.Password, u.Name, u.Email, u.Year, u.Month, u.Day, u.Sex)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return result
}

// isDuplicateUserID ユーザIDが既に登録されているかどうかを確認します。
func isDuplicateUserID(u User) bool {
	db := dbConnect()
	query := "SELECT count(*) FROM users WHERE userID = ?"

	cnt, err := db.Query(query, u.UserID)
	if err != nil {
		panic(err.Error())
	}
	// userID検索結果件数
	var Count Cnt
	for cnt.Next() {
		err := cnt.Scan(&Count.Count)
		if err != nil {
			panic(err.Error())
		}
	}
	defer db.Close()
	// ユーザIDが重複している場合、trueを返します。
	return Count.Count != 0
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
		return c.JSON(http.StatusOK, res)
	}
	// ユーザ登録メソッド実行
	result := registNewUser(*user)
	res.User = user
	res.Result = result
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
