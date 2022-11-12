package controller

import (
	"encoding/json"
	"hackathon/usecase"
	"log"
	"net/http"
)

func AffiliationSearch(w http.ResponseWriter, r *http.Request) {

	affiliations, statusCode := usecase.AffiliationSearch()
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(affiliations)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
