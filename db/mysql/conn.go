package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(192.168.159.128:3306)/claim?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		panic("Failed to connect to mysql2,err:" + err.Error())
	}
}

//DBConn：返回数据库链接对象
func DBConn() *sql.DB {
	return db
}
