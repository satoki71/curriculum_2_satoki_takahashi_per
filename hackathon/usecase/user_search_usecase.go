package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func UserSearch(name string) (users []model.UserResForHTTPGet, statusCode int) {
	if name == "" {
		log.Println("fail: name is empty")
		statusCode = 400
		return users, statusCode
	}

	rows, statusCode := dao.UserSearch(name)
	if statusCode != 0 {
		return users, statusCode
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}
	users = make([]model.UserResForHTTPGet, 0)
	for rows.Next() {
		var u model.UserResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.AffiliationId, &u.Points); err != nil {
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
		if err := allRows.Scan(&u.Id, &u.Name, &u.AffiliationId, &u.Points); err != nil {
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

func MemberUserSearch(affiliation string) (memberUsers []model.AllUserResForHTTPGet, statusCode int) {
	memberRows, statusCode := dao.MemberUserSearch(affiliation)
	if statusCode != 0 {
		return memberUsers, statusCode
	}

	memberUsers = make([]model.AllUserResForHTTPGet, 0)
	for memberRows.Next() {
		var u model.AllUserResForHTTPGet
		if err := memberRows.Scan(&u.Id, &u.Name, &u.AffiliationId, &u.Points); err != nil {
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
