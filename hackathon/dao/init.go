package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	_db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	//if err := _db.Ping(); err != nil {
	//	log.Fatalf("fail: _db.Ping, %v\n", err)
	//}
	db = _db

	//err := godotenv.Load(".envmysql")
	//if err != nil {
	//	log.Fatalf("fail: load envfile, %v\n", err)
	//}
	//// ①-1
	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlUserPwd := os.Getenv("MYSQL_PASSWORD")
	//mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	//
	//// ①-2
	//_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	//if err != nil {
	//	log.Fatalf("fail: sql.Open, %v\n", err)
	//}
	//// ①-3
	//if err := _db.Ping(); err != nil {
	//	log.Fatalf("fail: _db.Ping, %v\n", err)
	//}
	//db = _db

}
