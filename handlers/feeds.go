package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gratitude/db"
	"gratitude/models"
	"gratitude/shared"
	"net/http"
	"strconv"
)

type Feed struct {
	Friend models.FriendInfo `json:"info"`
	Notes  []models.Note     `json:"public_notes"`
}

func FeedsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["uuid"]

	offset := 0
	offsetQueryParam := r.URL.Query()["offset"]

	if offsetQueryParam != nil && len(offsetQueryParam[0]) > 0 {
		overrideOffset, err := strconv.Atoi(offsetQueryParam[0])
		if err != nil {
			errorResponse := shared.ErrorResponse{
				Message: err.Error(),
			}

			data, _ := json.Marshal(errorResponse)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(data)
		}
		offset = overrideOffset
	}

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
		notes := models.GetUserPublicNotes(client, friend.UID, offset)
		feeds = append(feeds, Feed{
			Notes:  notes,
			Friend: friend,
		})
	}

	data, _ := json.Marshal(feeds)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
