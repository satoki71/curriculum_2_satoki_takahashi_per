package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func GiveDelete(v model.GiveReqHTTPDelete) (statusCode int) {

	err := dao.GiveDelete(v)
	if err != nil {
		statusCode = 500
		return statusCode
	}

	return statusCode
}
