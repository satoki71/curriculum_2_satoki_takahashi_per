package controller

import (
	"encoding/json"
	"github.com/oklog/ulid"
	"hackathon/model"
	"hackathon/usecase"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GiveRegister(w http.ResponseWriter, r *http.Request) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id := ulid.MustNew(ms, entropy)

	jsonString, err := io.ReadAll(r.Body)
	//UserReqHTTPPost = io.ReadAll(r.Body)
	var v model.GiveReqHTTPPost
	if err := json.Unmarshal([]byte(jsonString), &v); err != nil {
		log.Printf("fail: json.Unmarshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statusCode := usecase.GiveRegister(id, v)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytesId, err := json.Marshal(model.UserIdReqPost{
		Id: id.String(),
	})
	if err != nil {
		log.Printf("fail: get id, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytesId)
	w.WriteHeader(http.StatusOK)

}
