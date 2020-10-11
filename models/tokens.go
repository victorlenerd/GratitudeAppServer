package models

import (
	"cloud.google.com/go/datastore"
	"context"
)

type UserToken struct {
	UserID string
	Token  string
}

func PutUserToken(client *datastore.Client, userID string, FCMToken string)  {
	ctx := context.Background()
	key := datastore.NameKey("UserToken", userID, nil)
	token := UserToken{
		UserID: userID,
		Token:  FCMToken,
	}
	_, err := client.Put(ctx, key, token)
	if err != nil {
		panic(err)
	}
}