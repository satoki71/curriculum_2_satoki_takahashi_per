package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func GiveSearch(name string) (allGiveUsers []model.AllGiveResForHTTPGet, statusCode int) {
	if name == "" {
		log.Println("fail: name is empty")
		statusCode = 400
		return allGiveUsers, statusCode
	}

	giveRows, statusCode := dao.GiveSearch(name)
	if statusCode != 0 {
		return allGiveUsers, statusCode
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}

	allGiveUsers = make([]model.AllGiveResForHTTPGet, 0)
	for giveRows.Next() {
		var u model.AllGiveResForHTTPGet
		if err := giveRows.Scan(&u.Id, &u.FromUserId, &u.Points, &u.Message, &u.ToUserId); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := giveRows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			statusCode = 500
			return allGiveUsers, statusCode
			//w.WriteHeader(http.StatusInternalServerError)
			//return
		}
		allGiveUsers = append(allGiveUsers, u)
	}
	return allGiveUsers, statusCode
}
