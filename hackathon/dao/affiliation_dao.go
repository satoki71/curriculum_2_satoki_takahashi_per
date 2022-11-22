package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"hackathon/model"
	"log"
)

func AffiliationSearch() (affiliationRows *sql.Rows, err error) {
	affiliationRows, err = db.Query("SELECT * FROM affiliation")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)

		return affiliationRows, err
	}
	return affiliationRows, err
}

func UAffiliationSearch(userId string) (rows *sql.Rows, err error) {
	rows, err = db.Query("SELECT * FROM affiliation WHERE id = (SELECT affiliationId FROM user WHERE userId = ?)", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)

		return rows, err
	}
	return rows, err
}

func AffiliationRegister(id ulid.ULID, v model.AffiliationReqForHTTPPost) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		return err
	}

	cmd := "INSERT INTO affiliation (id, name, number) VALUES (?, ?, ?)"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)

		return err
	}

	defer ins.Close()

	_, err = ins.Exec(id.String(), v.Name, 0)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)

		err = tx.Rollback()
		if err != nil {
			log.Printf("fail: tx.Rollback, %v\n", err)

		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)

	}
	return err
}
