package dao

import (
	"database/sql"
	"github.com/oklog/ulid"
	"hackathon/model"
	"log"
)

func UserSearch(userId string) (rows *sql.Rows, err error) {
	rows, err = db.Query("SELECT * FROM user WHERE userId = ?", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		//statusCode = 500
		return rows, err
	}
	return rows, err
}

func AllUserSearch() (allRows *sql.Rows, err error) {
	allRows, err = db.Query("SELECT * FROM user")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		//statusCode = 500
		return allRows, err
	}

	return allRows, err
}

func MemberUserSearch(name string) (memberRows *sql.Rows, err error) {
	memberRows, err = db.Query("SELECT * FROM user WHERE affiliationId = (SELECT id FROM affiliation WHERE name = ?)", name)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		//statusCode = 500
		return memberRows, err
	}
	return memberRows, err
}

func MateUserSearch(userId string) (memberRows *sql.Rows, err error) {
	memberRows, err = db.Query("SELECT * FROM user WHERE affiliationId = (SELECT affiliationId FROM user WHERE userId = ?) ORDER BY points DESC;", userId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		//statusCode = 500
		return memberRows, err
	}
	return memberRows, err
}

func UserRegister(id ulid.ULID, v model.UserReqHTTPPost) (err error) {

	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		//statusCode = 500
		return err
	}

	cmd := "INSERT INTO user (userId, name, affiliationId, points) VALUES (?, ?, (SELECT id FROM affiliation WHERE name = ?), ?)"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		//statusCode = 500
		return err
	}

	defer ins.Close()

	//subCmd := "SELECT id FROM affiliation WHERE name = ?, v.Affiliation"
	_, err = ins.Exec(id.String(), v.Name, v.Affiliation, 0)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		//statusCode = 500
		err := tx.Rollback()
		if err != nil {
			log.Printf("fail: tx.Rollback, %v\n", err)
			//statusCode = 500
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		//statusCode = 500
	}
	return err
}

func UserUpdate(userId string) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		//statusCode = 500
		return err
	}

	cmd := "UPDATE user SET user.points=(SELECT SUM(point.points) FROM point WHERE point.toUserId=?) WHERE user.userId=?"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		//statusCode = 500
		return err
	}

	defer ins.Close()

	_, err = ins.Exec(userId, userId)
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

func UserNameUpdate(v model.UserNameReqHTTPUpdate) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("fail: db.Begin, %v\n", err)
		//statusCode = 500
		return err
	}

	cmd := "UPDATE user SET name=? WHERE userId = ?"
	ins, err := tx.Prepare(cmd)
	if err != nil {
		log.Printf("fail: db.Prepare, %v\n", err)
		//statusCode = 500
		return err
	}

	defer ins.Close()

	_, err = ins.Exec(v.Name, v.UserId)
	if err != nil {
		log.Printf("fail: db.Exec, %v\n", err)
		//statusCode = 500
		err := tx.Rollback()
		if err != nil {
			log.Printf("fail: tx.Rollback, %v\n", err)
			//statusCode = 500
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("fail: tx.Commit, %v\n", err)
		//statusCode = 500
	}
	return err
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
