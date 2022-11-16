package controller

import (
	"hackathon/usecase"
	"net/http"
)

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId") // To be filled

	statusCode := usecase.UserUpdate(userId)
	if statusCode == 400 {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if statusCode == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	//w.Write()
	w.WriteHeader(http.StatusOK)

}
