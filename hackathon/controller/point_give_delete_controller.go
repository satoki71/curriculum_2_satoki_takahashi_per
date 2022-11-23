package controller

import (
	"encoding/json"
	"hackathon/model"
	"hackathon/usecase"
	"io"
	"log"
	"net/http"
)

func GiveDelete(w http.ResponseWriter, r *http.Request) {
	jsonString, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("fail: io.ReadAll, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var v model.GiveReqHTTPDelete
	if err := json.Unmarshal([]byte(jsonString), &v); err != nil {
		log.Printf("fail: json.Unmarshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statusCode := usecase.GiveDelete(v)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//w.Write()
	w.WriteHeader(http.StatusOK)
}
