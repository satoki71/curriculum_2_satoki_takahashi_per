package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"hackathon/model"
	"log"
)

func GiveSearch(userId string) (giveRows *sql.Rows, err error) {
	giveRows, err = db.Query("SELECT * FROM point WHERE fromUserId=?", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		//statusCode = 500
		return giveRows, err
	}
	return giveRows, err
}

func GiveRegister(id ulid.ULID, v model.GiveReqHTTPPost) (err error) {

	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		//statusCode = 500
		return err
	}

	cmd := "INSERT INTO point (id, fromUserId, points, message, toUserId) VALUES (?, ?, ?, ?, ?)"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		//statusCode = 500
		return err
	}

	defer ins.Close()

	//subCmd1 := "SELECT userId FROM user WHERE name = ?, v.FromName"
	//subCmd2 := "SELECT userId FROM user WHERE name = ?, v.ToName"
	_, err = ins.Exec(id.String(), v.FromUserId, v.Points, v.Message, v.ToUserId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		//statusCode = 500
		err1 := tx.Rollback()
		if err1 != nil {
			log.Printf("fail: tx.Rollback, %v\n", err1)
			//statusCode = 500
		}
		return err
	}
	err1 := tx.Commit()
	if err1 != nil {
		log.Printf("fail: tx.Commit, %v\n", err1)
		//statusCode = 500
	}
	return err
}

func GiveUpdate(v model.GiveReqHTTPPut) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)

		return err
	}

	cmd := "UPDATE point SET points=?, message=? WHERE id=?"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		return err
	}

	defer ins.Close()

	_, err = ins.Exec(v.Points, v.Message, v.Id)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)

		err1 := tx.Rollback()
		if err1 != nil {
			log.Printf("fail: tx.Rollback, %v\n", err1)

		}
		return err
	}
	err1 := tx.Commit()
	if err1 != nil {
		log.Printf("fail: tx.Commit, %v\n", err1)

	}
	return err
}

func GiveDelete(v model.GiveReqHTTPDelete) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)

		return err
	}
	_, err = tx.Exec("DELETE FROM point WHERE id=?", v.Id)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)

		err = tx.Rollback()
		if err != nil {
			log.Printf("fail: tx.Rollback, %v\n", err)

		}
		return err
	}
	err1 := tx.Commit()
	if err1 != nil {
		log.Printf("fail: tx.Commit, %v\n", err1)

	}
	return err
}
