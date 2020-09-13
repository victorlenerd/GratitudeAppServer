package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"time"
)

type Note struct {
	UUID       string    `json:"uuid"`
	Text       string    `json:"text"`
	IsPublic   bool      `json:"is_public"`
	OwnerID    string    `json:"owner_id"`
	Likes      int64     `json:"likes"`
	Views      int64     `json:"views"`
	CreateDate time.Time `json:"create_date"`
}

func CreateNewNote(client *datastore.Client, note *Note) {
	ctx := context.Background()
	key := datastore.NameKey("Note", note.UUID, nil)
	_, err := client.Put(ctx, key, note)
	if err != nil {
		panic(err)
	}
}

func GetUserNotes(client *datastore.Client, ownerID string) []Note {
	ctx := context.Background()
	query := datastore.NewQuery("Note").
		Filter("OwnerID =", ownerID)

	notes := []Note{}

	_, err := client.GetAll(ctx, query, &notes)
	if err != nil {
		panic(err)
	}

	return notes
}

func GetUserPublicNotes(client *datastore.Client, ownerID string, offset int) []Note {
	ctx := context.Background()
	query := datastore.NewQuery("Note").
		Filter("OwnerID =", ownerID).
		Filter("IsPublic =", true).
		Offset(offset).
		Limit(20)

	notes := []Note{}

	_, err := client.GetAll(ctx, query, &notes)
	if err != nil {
		panic(err)
	}

	return notes
}

func DeleteNote(client *datastore.Client, noteID string) {
	ctx := context.Background()
	key := datastore.NameKey("Note", noteID, nil)
	err := client.Delete(ctx, key)

	if err != nil {
		fmt.Println(err)
	}
}
