package infrastructure

import "database/sql"

// DbConnect Mysqlへ接続し、dbインスタンスを返却します。
func DbConnect() *sql.DB {
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

// Transact トランザクションを管理するラッパ関数
func Transact(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return err
}
