package controller

import (
	"encoding/json"
	"hackathon/usecase"
	"log"
	"net/http"
)

func UserSearch(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // To be filled

	users, statusCode := usecase.UserSearch(name)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func AllUserSearch(w http.ResponseWriter, r *http.Request) {

	allUsers, statusCode := usecase.AllUserSearch()
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(allUsers)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func MemberUserSearch(w http.ResponseWriter, r *http.Request) {
	affiliation := r.URL.Query().Get("affiliation") // To be filled

	//usecase関数が必要(userSearch参照)
	memberUsers, statusCode := usecase.MemberUserSearch(affiliation)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(memberUsers)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
