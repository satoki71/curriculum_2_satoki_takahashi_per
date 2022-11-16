package usecase

import (
	"hackathon/dao"
	"log"
)

func UserUpdate(userId string) (statusCode int) {
	if userId == "" {
		log.Println("fail: userId is empty")
		statusCode = 400
		return statusCode
	}

	statusCode = dao.UserUpdate(userId)

	return statusCode
}
