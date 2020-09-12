package db

import (
	"cloud.google.com/go/datastore"
	"context"
)

var cachedDataStoreClient *datastore.Client = nil

func Init(ctx context.Context) {
	client, err := datastore.NewClient(ctx, "gratitude-app-server")

	if err != nil {
		panic(err)
	}

	cachedDataStoreClient = client
}

func GetClient() *datastore.Client {
	if cachedDataStoreClient == nil {
		panic("data store db has not been initialized")
	}

	return cachedDataStoreClient
}
