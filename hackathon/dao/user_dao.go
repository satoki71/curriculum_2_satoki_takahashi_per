package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"hackathon/model"
	"log"
)

func UserSearch(name string) (rows *sql.Rows, statusCode int) {
	rows, err := db.Query("SELECT * FROM user WHERE name = ?", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return rows, statusCode
	}
	return rows, statusCode
}

func AllUserSearch() (allRows *sql.Rows, statusCode int) {
	allRows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return allRows, statusCode
	}
	return allRows, statusCode
}

func MemberUserSearch(name string) (memberRows *sql.Rows, statusCode int) {
	memberRows, err := db.Query("SELECT * FROM user WHERE affiliationId = (SELICT id FROM affiliation WHERE name = ?)", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		statusCode = 500
		return memberRows, statusCode
	}
	return memberRows, statusCode
}

func UserRegister(id ulid.ULID, v model.UserReqHTTPPost) (statusCode int) {

	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}

	cmd := "INSERT INTO user (userId, name, affiliationId, points) VALUES (?, ?, (SELECT id FROM affiliation WHERE name = ?), ?)"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		statusCode = 500
		return statusCode
	}

	defer ins.Close()

	//subCmd := "SELECT id FROM affiliation WHERE name = ?, v.Affiliation"
	_, err = ins.Exec(id.String(), v.Name, v.Affiliation, 0)
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

func UserUpdate(name string) (statusCode int) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		statusCode = 500
		return statusCode
	}

	cmd := "UPDATE user SET points=(SELECT SUM(points) FROM point WHERE toUserId=(SELECT userId FROM user WHERE name=?)) WHERE name=?"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		statusCode = 500
		return statusCode
	}

	defer ins.Close()

	_, err = ins.Exec(name, name)
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

// ③ Ctrl+CでHTTPサーバー停止時にDBをクローズする

//func CloseDBWithSysCall() {
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
//	go func() {
//		s := <-sig
//		log.Printf("received syscall, %v", s)
//
//		if err := db.Close(); err != nil {
//			log.Fatal(err)
//		}
//		log.Printf("success: db.Close()")
//		os.Exit(0)
//	}()
//}
