package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func AffiliationSearch() (affiliations []model.AffiliationResForHTTPGet, statusCode int) {
	allRows, err := dao.AffiliationSearch()
	if err != nil {
		statusCode = 500
		return affiliations, statusCode
	}

	affiliations = make([]model.AffiliationResForHTTPGet, 0)
	for allRows.Next() {
		var u model.AffiliationResForHTTPGet
		if err := allRows.Scan(&u.Id, &u.Name, &u.Number); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := allRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return affiliations, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		affiliations = append(affiliations, u)
	}
	return affiliations, statusCode
}

func UAffiliationSearch(userId string) (affiliation []model.AffiliationResForHTTPGet, statusCode int) {
	if userId == "" {
		log.Println("fail: userId is empty")
		statusCode = 400
		return affiliation, statusCode
	}

	rows, err := dao.UAffiliationSearch(userId)
	if err != nil {
		statusCode = 500
		return affiliation, statusCode
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}
	affiliation = make([]model.AffiliationResForHTTPGet, 0)
	for rows.Next() {
		var u model.AffiliationResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Number); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return affiliation, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		affiliation = append(affiliation, u)
	}
	return affiliation, statusCode
}
