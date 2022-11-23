package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func GiveUpdate(v model.GiveReqHTTPPut) (statusCode int) {

	err := dao.GiveUpdate(v)
	if err != nil {
		statusCode = 500
		return statusCode
	}

	return statusCode
}
