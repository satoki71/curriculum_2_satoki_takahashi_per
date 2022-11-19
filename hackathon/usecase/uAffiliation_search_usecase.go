package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func UAffiliationSearch(userId string) (affiliation []model.AffiliationResForHTTPGet, statusCode int) {
	if userId == "" {
		log.Println("fail: userId is empty")
		statusCode = 400
		return affiliation, statusCode
	}

	rows, statusCode := dao.UAffiliationSearch(userId)
	if statusCode != 0 {
		return affiliation, statusCode
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}
	affiliation = make([]model.AffiliationResForHTTPGet, 0)
	for rows.Next() {
		var u model.AffiliationResForHTTPGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Name, &u.Number); err != nil {
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
