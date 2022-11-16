package usecase

import (
	"github.com/oklog/ulid"
	"hackathon/dao"
	"hackathon/model"
)

func GiveRegister(id ulid.ULID, v model.GiveReqHTTPPost) (statusCode int) {

	statusCode = dao.GiveRegister(id, v)
	//if statusCode == 500 {
	//	return statusCode
	//}
	//statusCode = dao.UserUpdate(v.ToUserId)

	return statusCode
}
