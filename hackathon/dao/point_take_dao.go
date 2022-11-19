package dao

import (
	"database/sql"
	"log"
)

func TakeSearch(userId string) (takeRows *sql.Rows, err error) {
	takeRows, err = db.Query("SELECT * FROM point WHERE toUserId=?", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		//statusCode = 500
		return takeRows, err
	}
	return takeRows, err
}
