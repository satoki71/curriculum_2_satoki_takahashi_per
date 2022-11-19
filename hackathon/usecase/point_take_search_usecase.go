package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func TakeSearch(userid string) (allTakeUsers []model.AllTakeResForHTTPGet, statusCode int) {
	if userid == "" {
		log.Println("fail: userid is empty")
		statusCode = 400
		return allTakeUsers, statusCode
	}

	takeRows, err := dao.TakeSearch(userid)
	if err != nil {
		statusCode = 500
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
