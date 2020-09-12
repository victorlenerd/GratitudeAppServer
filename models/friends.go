package models

import (
	"cloud.google.com/go/datastore"
	"time"
)

type FriendStatus int

const (
	Approved FriendStatus = iota + 1
	Pending
	Declined
)

type Friend struct {
	UUID			string
	OwnerID 		string
	UserID          string
	Status 			FriendStatus
	Created         time.Time
}

type FriendModel Friend

func (* FriendModel) CreateNewPendingRequest(client *datastore.Client) {

}