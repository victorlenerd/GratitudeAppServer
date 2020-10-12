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
	_, err := client.Put(ctx, key, &token)
	if err != nil {
		panic(err)
	}
}

func GetUsersToken(client *datastore.Client, userIDs []string) []string {
	ctx := context.Background()
	keys :=	[]*datastore.Key{}

	for _, uuid := range userIDs {
		keys = append(keys, datastore.NameKey("UserToken", uuid, nil))
	}

	userTokens := []UserToken{}

	err := client.GetMulti(ctx, keys, userTokens)
	if err != nil {
		panic(err)
	}

	tokens := []string{}

	for _, userToken := range userTokens {
		tokens = append(tokens, userToken.Token)
	}

	return tokens
}