package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func GiveUpdate(v model.GiveReqHTTPPut) (statusCode int) {

	statusCode = dao.GiveUpdate(v)

	return statusCode
}
