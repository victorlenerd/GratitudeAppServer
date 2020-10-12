package shared

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"gratitude/db"
	"gratitude/models"
)

var FirebaseAppOpts = option.WithCredentialsFile("gratitude-8563a-firebase-adminsdk-om7ze-7f87952725.json")

func SendFeedsNotifications(userID string) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, FirebaseAppOpts)
	if err != nil {
		panic(err)
	}

	messagingClient, err := app.Messaging(ctx)
	if err != nil {
		panic(err)
	}

	dbClient := db.GetClient()
	friends := models.GetAllFriends(dbClient, userID)
	userIDs := make([]string, len(friends))

	for _, friend := range friends {
		if friend.Request.Status == "3" {
			userIDs = append(userIDs, friend.Info.UID)
		}
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}


	userInfo, err := authClient.GetUser(ctx, userID)
	if err != nil {
		panic(err)
	}

	tokens := models.GetUsersToken(dbClient, userIDs)
	notification := &messaging.Notification{
		Title: "New Feed Note",
		Body: 	userInfo.DisplayName+" shared a note",
	}
	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Notification: notification,
	}
	_, err = messagingClient.SendMulticast(ctx, message)
	if err != nil {
		panic(err)
	}
}
