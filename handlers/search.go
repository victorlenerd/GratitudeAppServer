package handlers

import (
	"encoding/json"
	"gratitude/models"
	"gratitude/shared"
	"net/http"
)

func SearchFriendHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query()["email"]

	if email == nil || len(email[0]) < 1 {
		errorResponse := shared.ErrorResponse{
			Message: "email is required params",
		}

		data, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	user := models.SearchForFriendByEmail(email[0])

	data, _ := json.Marshal(user)

	if user != nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write(data)
}