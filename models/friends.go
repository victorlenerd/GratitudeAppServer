package models

import (
	"cloud.google.com/go/datastore"
	"context"
	firebase "google.golang.org/api/firebase/v1beta1"
	"log"
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
	UserID         	string
	OwnerID         string
	Status 			FriendStatus
	CreatedDate     time.Time
}

func CreateFriends(client *datastore.Client, friend Friend) {
	ctx := context.Background()
	key := datastore.NameKey("Friend", friend.UUID, nil)
	_, err := client.Put(ctx, key, &friend)
	if err != nil {
		panic(err)
	}
}

func GetAllFriends(client *datastore.Client, ownerID string) []Friend {
	ctx := context.Background()

	ownerBasedQuery := datastore.NewQuery("Friend").
		Filter("OwnerID =", ownerID)

	ownerBasedQueryFriends := []Friend{}
	_, err := client.GetAll(ctx, ownerBasedQuery, &ownerBasedQueryFriends)
	if err != nil {
		panic(err)
	}

	if len(ownerBasedQueryFriends) > 0 {
		return ownerBasedQueryFriends
	}

	userBasedQuery := datastore.NewQuery("Friend").
		Filter("UserID =", ownerID)

	userBasedQueryFriends := []Friend{}
	_, err = client.GetAll(ctx, userBasedQuery, &userBasedQueryFriends)
	if err != nil {
		panic(err)
	}

	return userBasedQueryFriends
}

func SearchForFriendByEmail() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

}