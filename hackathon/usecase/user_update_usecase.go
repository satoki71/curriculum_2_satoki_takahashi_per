package usecase

import (
	"hackathon/dao"
	"log"
)

func UserUpdate(name string) (statusCode int) {
	if name == "" {
		log.Println("fail: name is empty")
		statusCode = 400
		return statusCode
	}

	statusCode = dao.UserUpdate(name)

	return statusCode
}
