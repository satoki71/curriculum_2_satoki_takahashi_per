package usecase

import (
	"github.com/oklog/ulid"
	"hackathon/dao"
	"hackathon/model"
)

func GiveRegister(id ulid.ULID, v model.GiveReqHTTPPost) (statusCode int) {

	statusCode = dao.GiveRegister(id, v)

	statusCode = dao.UserUpdate(v.ToUserId)

	return
}
