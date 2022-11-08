package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func TakeSearch(name string) (allTakeUsers []model.AllTakeResForHTTPGet, statusCode int) {
	if name == "" {
		log.Println("fail: name is empty")
		statusCode = 400
		return allTakeUsers, statusCode
	}

	takeRows, statusCode := dao.TakeSearch(name)
	if statusCode != 0 {
		return allTakeUsers, statusCode
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}
	allTakeUsers = make([]model.AllTakeResForHTTPGet, 0)
	for takeRows.Next() {
		var u model.AllTakeResForHTTPGet
		if err := takeRows.Scan(&u.Id, &u.FromUserId, &u.Points, &u.Message, &u.ToUserId); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			if err := takeRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return allTakeUsers, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		allTakeUsers = append(allTakeUsers, u)
	}
	return allTakeUsers, statusCode
}
