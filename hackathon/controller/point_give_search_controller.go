package controller

import (
	"encoding/json"
	"hackathon/usecase"
	"log"
	"net/http"
)

func GiveSearch(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // To be filled

	//usecase関数が必要(userSearch参照)
	allGiveUsers, statusCode := usecase.GiveSearch(name)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(allGiveUsers)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
