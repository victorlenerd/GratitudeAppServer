package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gratitude/db"
	"gratitude/models"
	"gratitude/shared"
	"io/ioutil"
	"net/http"
	"time"
)

type PutFriendRequestBody struct {
	UUID    string `json:"uuid"`
	UserID  string `json:"user_id"`
	OwnerID string `json:"owner_id"`
	Status 	string `json:"status"`
}

func PutFriendHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse := shared.ErrorResponse{
			Message: err.Error(),
		}

		data, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	body := PutFriendRequestBody{}

	err = json.Unmarshal(data, &body)
	if err != nil {
		errorResponse := shared.ErrorResponse{
			Message: err.Error(),
		}

		data, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	client := db.GetClient()

	friendModel := models.FriendRequest{
		UUID: body.UUID,
		OwnerID: body.OwnerID,
		UserID: body.UserID,
		CreatedDate: time.Now(),
		Status: body.Status,
	}

	models.CreateFriends(client, friendModel)

	if friendModel.Status == "1" {
		shared.SendPendingFriendRequestNotification(friendModel.OwnerID, friendModel.UserID)
	}

	data, _ = json.Marshal(friendModel)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func GetAllFriendHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ownerID := params["ownerID"]
	if len(ownerID) < 1 {
		errorResponse := shared.ErrorResponse{
			Message: "ownerID is required params",
		}

		data, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	client := db.GetClient()
	friends := models.GetAllFriends(client, ownerID)

	data, _ := json.Marshal(friends)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteFriendHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	uuid := params["uuid"]
	if len(uuid) < 1 {
		errorResponse := shared.ErrorResponse{
			Message: "uuid is required params",
		}

		data, _ := json.Marshal(errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	client := db.GetClient()
	models.DeleteFriend(client, uuid)

	w.WriteHeader(http.StatusNoContent)
	w.Write(nil)
}