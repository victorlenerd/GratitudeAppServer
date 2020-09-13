package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gratitude/db"
	"gratitude/models"
	"gratitude/shared"
	"net/http"
)

type Feed struct {
	Friend models.FriendInfo `json:"info"`
	Notes  []models.Note     `json:"public_notes"`
}

func FeedsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)
	uuid := params["uuid"]

	if len(uuid) < 0 {
		errorResponse := shared.ErrorResponse{
			Message: "uuid is required params",
		}

		data, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	client := db.GetClient()
	friends := models.GetAllFriends(client, uuid)

	feeds := []Feed{}

	for _, friend := range friends {
		notes := models.GetUserPublicNotes(client, friend.UID)
		feeds = append(feeds, Feed{
			Notes:  notes,
			Friend: friend,
		})
	}

	data, _ := json.Marshal(feeds)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
