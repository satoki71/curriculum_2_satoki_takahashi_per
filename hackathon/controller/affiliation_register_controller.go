package controller

import (
	"hackathon/usecase"
	"net/http"
)

func AffiliationRegister(w http.ResponseWriter, r *http.Request) {
	//entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	//ms := ulid.Timestamp(time.Now())
	//id := ulid.MustNew(ms, entropy)
	//add id

	affiliation := r.URL.Query().Get("name") // To be filled
	statusCode := usecase.AffiliationRegister(affiliation)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//bytesId, err := json.Marshal(model.UserIdReqPost{
	//	Id: id.String(),
	//})
	//if err != nil {
	//	log.Printf("fail: get id, %v\n", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")
	//w.Write(bytesId)
	w.WriteHeader(http.StatusOK)
}
