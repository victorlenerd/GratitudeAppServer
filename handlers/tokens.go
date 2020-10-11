package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gratitude/db"
	"gratitude/models"
	"gratitude/shared"
	"io/ioutil"
	"net/http"
)

type PutUserTokenRequestBody struct {
	FCMToken    string `json:"fcm_token"`
}

func PutUserToken(w http.ResponseWriter, r *http.Request) {
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

	body := PutUserTokenRequestBody{}

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
	models.PutUserToken(client, ownerID, body.FCMToken)
	w.WriteHeader(http.StatusOK)

	if err != nil {
		panic(err)
	}

	w.Write(data)
}
