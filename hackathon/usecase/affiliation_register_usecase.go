package usecase

import (
	"github.com/oklog/ulid"
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func AffiliationRegister(id ulid.ULID, v model.AffiliationReqForHTTPPost) (statusCode int) {
	if v.Name == "" || len(v.Name) >= 50 {
		log.Println("fail: name is empty or too long")
		statusCode = 400
		return statusCode
	}

	err := dao.AffiliationRegister(id, v)
	if err != nil {
		statusCode = 500
	}

	return statusCode
}
