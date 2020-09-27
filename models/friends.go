package models

import (
	"cloud.google.com/go/datastore"
	"context"
	firebase "firebase.google.com/go"
	"time"
)

type FriendStatus int

const (
	Approved FriendStatus = iota + 1
	Pending
	Declined
)

type FriendRequest struct {
	UUID        string
	UserID      string
	OwnerID     string
	Status      FriendStatus
	CreatedDate time.Time
}

type FriendInfo struct {
	UUID        string `json:"uuid"`
	UID         string `json:"uid"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type FriendContainer struct {
	friendInfo    FriendInfo    `json:"friend_info"`
	friendRequest FriendRequest `json:"friend_request"`
}

func CreateFriends(client *datastore.Client, friend FriendRequest) {
	ctx := context.Background()
	key := datastore.NameKey("FriendRequest", friend.UUID, nil)
	_, err := client.Put(ctx, key, &friend)
	if err != nil {
		panic(err)
	}
}

func GetAllFriends(client *datastore.Client, ownerID string) []FriendInfo {
	ctx := context.Background()

	ownerBasedQuery := datastore.NewQuery("FriendRequest").
		Filter("OwnerID =", ownerID)

	ownerBasedQueryFriends := []FriendRequest{}
	_, err := client.GetAll(ctx, ownerBasedQuery, &ownerBasedQueryFriends)
	if err != nil {
		panic(err)
	}

	if len(ownerBasedQueryFriends) > 0 {
		uids := [][]string{}

		for _, userInfo := range ownerBasedQueryFriends {
			uids = append(uids, []string{userInfo.UUID, userInfo.UserID})
		}

		return getUserFriends(uids)
	}

	userBasedQuery := datastore.NewQuery("FriendRequest").
		Filter("UserID =", ownerID)

	userBasedQueryFriends := []FriendRequest{}
	_, err = client.GetAll(ctx, userBasedQuery, &userBasedQueryFriends)
	if err != nil {
		panic(err)
	}

	uids := [][]string{}
	for _, userInfo := range userBasedQueryFriends {
		uids = append(uids, []string{userInfo.UUID, userInfo.UserID})
	}

	return getUserFriends(uids)
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

func getUserFriends(uids [][]string) []FriendInfo {
	friendsInfo := []FriendInfo{}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	for _, uid := range uids {
		userRecord, _ := authClient.GetUser(ctx, uid[1])

		if userRecord != nil {
			friendsInfo = append(friendsInfo, FriendInfo{
				UUID:        uid[0],
				UID:         userRecord.UID,
				DisplayName: userRecord.DisplayName,
				Email:       userRecord.Email,
			})
		}

	}

	return friendsInfo
}
