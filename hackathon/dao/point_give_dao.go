package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"hackathon/model"
	"log"
)

func GiveSearch(userId string) (giveRows *sql.Rows, statusCode int) {
	giveRows, err := db.Query("SELECT * FROM point WHERE fromUserId=?", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return giveRows, statusCode
	}
	return giveRows, statusCode
}

func GiveRegister(id ulid.ULID, v model.GiveReqHTTPPost) (statusCode int) {

	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}

	cmd := "INSERT INTO point (id, fromUserId, points, message, toUserId) VALUES (?, (SELECT userId FROM user WHERE name = ?), ?, ?, (SELECT userId FROM user WHERE name = ?))"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		statusCode = 500
		return statusCode
	}

	defer ins.Close()

	//subCmd1 := "SELECT userId FROM user WHERE name = ?, v.FromName"
	//subCmd2 := "SELECT userId FROM user WHERE name = ?, v.ToName"
	_, err = ins.Exec(id.String(), v.FromName, v.Points, v.Message, v.ToName)
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

func GiveUpdate(v model.GiveReqHTTPPut) (statusCode int) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}

	cmd := "UPDATE point SET points=?, message=? WHERE id=?"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		statusCode = 500
		return statusCode
	}

	defer ins.Close()

	_, err = ins.Exec(v.Points, v.Message, v.Id)
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

func GiveDelete(v model.GiveReqHTTPDelete) (statusCode int) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}
	_, err = tx.Exec("DELETE FROM point WHERE id=?", v.Id)
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
