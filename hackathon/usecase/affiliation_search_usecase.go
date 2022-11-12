package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func AffiliationSearch() (affiliations []model.AffiliationResForHTTPGet, statusCode int) {
	allRows, statusCode := dao.AffiliationSearch()
	if statusCode != 0 {
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
