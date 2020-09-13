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

type PutNoteRequestBody struct {
	UUID       string    `json:"uuid"`
	Text       string    `json:"text"`
	IsPublic   bool      `json:"is_public"`
	OwnerID    string    `json:"owner_id"`
	Likes      int64     `json:"likes"`
	Views      int64     `json:"views"`
}

func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	errorResponse := shared.ErrorResponse{}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorResponse.Message = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	}

	noteBody := PutNoteRequestBody{}

	err = json.Unmarshal(body, &noteBody)
	if err != nil {
		panic(err)
	}

	note := models.Note{
		UUID:       noteBody.UUID,
		Text:       noteBody.Text,
		IsPublic:   noteBody.IsPublic,
		OwnerID:    noteBody.OwnerID,
		Likes:      noteBody.Likes,
		Views:      noteBody.Views,
		CreateDate: time.Now(),
	}
	
	client := db.GetClient()
	models.CreateNewNote(client, &note)

	if err != nil {
		errorResponse.Message = err.Error()
		data, _ := json.Marshal(&errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(nil)
}

func GetAllNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	if ownerID, ok := r.URL.Query()["ownerID"]; !ok || len(ownerID[0]) < 1 {
		errorResponse := shared.ErrorResponse{
			Message: "owner id is required",
		}
		data, _ := json.Marshal(&errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
	} else {
		client := db.GetClient()
		notes := models.GetUserNotes(client, ownerID[0])
		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(&notes)

		if err != nil {
			panic(err)
		}

		w.Write(data)
	}
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(r)

	noteUUID := params["uuid"]

	if len(noteUUID) < 1 {
		errorResponse := shared.ErrorResponse{
			Message: "note uui id is required",
		}
		data, _ := json.Marshal(&errorResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
	} else {
		client := db.GetClient()
		models.DeleteNote(client, noteUUID)
		w.WriteHeader(http.StatusNoContent)
		w.Write(nil)
	}
}