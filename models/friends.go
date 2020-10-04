package models

import (
	"cloud.google.com/go/datastore"
	"context"
	firebase "firebase.google.com/go"
	"time"
)

type FriendStatus int

type FriendRequest struct {
	UUID        string       `json:"uuid"`
	UserID      string       `json:"user_id"`
	OwnerID     string       `json:"owner_id"`
	Status      string 		 `json:"status"`
	CreatedDate time.Time    `json:"created_date"`
}

type FriendInfo struct {
	UUID        string `json:"uuid"`
	UID         string `json:"uid"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type FriendContainer struct {
	Request FriendRequest `json:"request"`
	Info    FriendInfo    `json:"info"`
}

func CreateFriends(client *datastore.Client, friend FriendRequest) {
	ctx := context.Background()
	key := datastore.NameKey("FriendRequest", friend.UUID, nil)
	_, err := client.Put(ctx, key, &friend)
	if err != nil {
		panic(err)
	}
}

func GetAllFriends(client *datastore.Client, ownerID string) []FriendContainer {
	ctx := context.Background()

	ownerBasedQuery := datastore.NewQuery("FriendRequest").
		Filter("OwnerID =", ownerID)

	requests := []FriendRequest{}

	_, err := client.GetAll(ctx, ownerBasedQuery, &requests)
	if err != nil {
		panic(err)
	}

	UIDS := [][]string{}

	if len(requests) > 0 {
		for _, userInfo := range requests {
			UIDS = append(UIDS, []string{userInfo.UUID, userInfo.UserID})
		}
	} else {
		userBasedQuery := datastore.NewQuery("FriendRequest").
			Filter("UserID =", ownerID)

		_, err = client.GetAll(ctx, userBasedQuery, &requests)
		if err != nil {
			panic(err)
		}

		for _, userInfo := range requests {
			UIDS = append(UIDS, []string{userInfo.UUID, userInfo.OwnerID})
		}
	}

	return getUserFriends(UIDS, requests)
}

func DeleteFriend(client *datastore.Client, uuid string) {
	ctx := context.Background()
	key := datastore.NameKey("FriendRequest", uuid, nil)
	err := client.Delete(ctx, key)
	if err != nil {
		panic(err)
	}
}

func SearchForFriendByEmail(email string) *FriendInfo {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	userRecord, _ := authClient.GetUserByEmail(ctx, email)
	if userRecord != nil {
		return &FriendInfo{
			UUID:        userRecord.UID,
			UID:         userRecord.UID,
			DisplayName: userRecord.DisplayName,
			Email:       userRecord.Email,
		}
	}

	return nil
}

func getUserFriends(uids [][]string, requests []FriendRequest) []FriendContainer {
	friends := []FriendContainer{}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	for i, uid := range uids {
		userRecord, _ := authClient.GetUser(ctx, uid[1])

		if userRecord != nil {
			friends = append(friends, FriendContainer{
				Info: FriendInfo{
					UUID:        uid[0],
					UID:         userRecord.UID,
					DisplayName: userRecord.DisplayName,
					Email:       userRecord.Email,
				},
				Request: requests[i],
			})
		}

	}

	return friends
}
