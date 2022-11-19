package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"log"
)

func AffiliationSearch() (affiliationRows *sql.Rows, statusCode int) {
	affiliationRows, err := db.Query("SELECT * FROM affiliation")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return affiliationRows, statusCode
	}
	return affiliationRows, statusCode
}

func UAffiliationSearch(userId string) (rows *sql.Rows, statusCode int) {
	rows, err := db.Query("SELECT * FROM affiliation WHERE id = (SELECT affiliationId FROM user WHERE userId = ?)", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return rows, statusCode
	}
	return rows, statusCode
}

func AffiliationRegister(id ulid.ULID, affiliation string) (statusCode int) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}

	cmd := "INSERT INTO affiliation (id, name) VALUES (?, ?)"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		statusCode = 500
		return statusCode
	}

	defer ins.Close()

	_, err = ins.Exec(id.String(), affiliation)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		statusCode = 500
		err1 := tx.Rollback()
		if err1 != nil {
			log.Printf("fail: tx.Rollback, %v\n", err1)
			statusCode = 500
		}
		return statusCode
	}
	err1 := tx.Commit()
	if err1 != nil {
		log.Printf("fail: tx.Commit, %v\n", err1)
		statusCode = 500
	}
	return statusCode
}
