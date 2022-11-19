package usecase

import (
	"github.com/oklog/ulid"
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func UserRegister(id ulid.ULID, v model.UserReqHTTPPost) (statusCode int) {
	//statusCode = 0
	if v.Name == "" || len(v.Name) >= 50 {
		log.Println("fail: name is empty or too long")
		statusCode = 400
		return statusCode
	}

	//if v.Age < 20 || v.Age > 80 {
	//	log.Println("fail: age is too short or too long")
	//	statusCode = 400
	//	return statusCode
	//}

	err := dao.UserRegister(id, v)
	if err != nil {
		statusCode = 500
		return statusCode
	}

	return statusCode
}
