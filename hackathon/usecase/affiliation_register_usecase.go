package usecase

import (
	"hackathon/dao"
	"log"
)

func AffiliationRegister(affiliation string) (statusCode int) {
	if affiliation == "" || len(affiliation) >= 50 {
		log.Println("fail: name is empty or too long")
		statusCode = 400
		return statusCode
	}

	statusCode = dao.AffiliationRegister(affiliation)

	return statusCode
}
