package dao

import (
	"database/sql"
	"log"
)

func TakeSearch(name string) (takeRows *sql.Rows, statusCode int) {
	takeRows, err := db.Query("SELECT * FROM point WHERE toUserId=(SELICT userId FROM user WHERE name = ?)", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return takeRows, statusCode
	}
	return takeRows, statusCode
}
