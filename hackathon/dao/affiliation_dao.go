package dao

import "log"

func AffiliationRegister(affiliation string) (statusCode int) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}

	_, err = tx.Exec("INSERT INTO affiliation (name) VALUES (?)", affiliation)
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
