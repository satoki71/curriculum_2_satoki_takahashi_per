package usecase

import (
	"hackathon/dao"
	"hackathon/model"
)

func GiveDelete(v model.GiveReqHTTPDelete) (statusCode int) {

	statusCode = dao.GiveDelete(v)

	return statusCode
}
