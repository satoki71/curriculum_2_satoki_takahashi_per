package usecase

import (
	"hackathon/dao"
	"hackathon/model"
	"log"
)

func UserUpdate(userId string) (statusCode int) {
	if userId == "" {
		log.Println("fail: userId is empty")
		statusCode = 400
		return statusCode
	}

	err := dao.UserUpdate(userId)
	if err != nil {
		statusCode = 500
		return statusCode
	}

	return statusCode
}

func UserNameUpdate(v model.UserNameReqHTTPUpdate) (statusCode int) {
	if v.Name == "" || len(v.Name) >= 50 {
		log.Println("fail: name is empty or too long")
		statusCode = 400
		return statusCode
	}

	err := dao.UserNameUpdate(v)
	if err != nil {
		statusCode = 500
		return statusCode
	}

	return statusCode
}
