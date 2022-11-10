package usecase

import (
	"github.com/oklog/ulid"
	"hackathon/dao"
	"log"
)

func AffiliationRegister(id ulid.ULID, affiliation string) (statusCode int) {
	if affiliation == "" || len(affiliation) >= 50 {
		log.Println("fail: name is empty or too long")
		statusCode = 400
		return statusCode
	}

	statusCode = dao.AffiliationRegister(id, affiliation)

	return statusCode
}
