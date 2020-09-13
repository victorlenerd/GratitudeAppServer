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
	Status 	models.FriendStatus
}

func PutFriendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

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

	friendModel := models.Friend{
		UUID: body.UUID,
		OwnerID: body.OwnerID,
		UserID: body.UserID,
		CreatedDate: time.Now(),
		Status: body.Status,
	}

	models.CreateFriends(client, friendModel)

	data, _ = json.Marshal(friendModel)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func DeleteFriendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

}

func GetOneFriendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")


}

func GetAllFriendHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

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
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
