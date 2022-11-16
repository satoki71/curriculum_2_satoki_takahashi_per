package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func UserSearch(userId string) (users []model.UserResForHTTPGet, statusCode int) {
	if userId == "" {
		log.Println("fail: userId is empty")
		statusCode = 400
		return users, statusCode
	}

	rows, statusCode := dao.UserSearch(userId)
	if statusCode != 0 {
		return users, statusCode
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}
	users = make([]model.UserResForHTTPGet, 0)
	for rows.Next() {
		var u model.UserResForHTTPGet
		if err := rows.Scan(&u.UserId, &u.Name, &u.AffiliationId, &u.Points); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return users, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		users = append(users, u)
	}
	return users, statusCode
}

func AllUserSearch() (allUsers []model.AllUserResForHTTPGet, statusCode int) {
	allRows, statusCode := dao.AllUserSearch()
	if statusCode != 0 {
		return allUsers, statusCode
	}

	allUsers = make([]model.AllUserResForHTTPGet, 0)
	for allRows.Next() {
		var u model.AllUserResForHTTPGet
		if err := allRows.Scan(&u.UserId, &u.Name, &u.AffiliationId, &u.Points); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := allRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return allUsers, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		allUsers = append(allUsers, u)
	}
	return allUsers, statusCode
}

func MemberUserSearch(name string) (memberUsers []model.MemberUserResForHTTPGet, statusCode int) {
	if name == "" {
		log.Println("fail: name is empty")
		statusCode = 400
		return memberUsers, statusCode
	}

	memberRows, statusCode := dao.MemberUserSearch(name)
	if statusCode != 0 {
		return memberUsers, statusCode
	}

	memberUsers = make([]model.MemberUserResForHTTPGet, 0)
	for memberRows.Next() {
		var u model.MemberUserResForHTTPGet
		if err := memberRows.Scan(&u.UserId, &u.Name, &u.AffiliationId, &u.Points); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := memberRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return memberUsers, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		memberUsers = append(memberUsers, u)
	}
	return memberUsers, statusCode
}

func MateUserSearch(userId string) (memberUsers []model.MemberUserResForHTTPGet, statusCode int) {
	if userId == "" {
		log.Println("fail: userId is empty")
		statusCode = 400
		return memberUsers, statusCode
	}

	memberRows, statusCode := dao.MateUserSearch(userId)
	if statusCode != 0 {
		return memberUsers, statusCode
	}

	memberUsers = make([]model.MemberUserResForHTTPGet, 0)
	for memberRows.Next() {
		var u model.MemberUserResForHTTPGet
		if err := memberRows.Scan(&u.UserId, &u.Name, &u.AffiliationId, &u.Points); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := memberRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return memberUsers, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		memberUsers = append(memberUsers, u)
	}
	return memberUsers, statusCode
}
