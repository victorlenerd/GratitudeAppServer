package main

import (
	"context"
	"github.com/gorilla/mux"
	"gratitude/db"
	"gratitude/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")

	// Initialize Data store Client
	ctx := context.Background()
	db.Init(ctx)

	// Feeds endpoint
	router.HandleFunc("/feeds", handlers.FeedsHandler).Methods(http.MethodGet)

	// Notes endpoint

	router.HandleFunc("/notes", handlers.GetAllNoteHandler).Methods(http.MethodGet)
	router.HandleFunc("/notes", handlers.PutNoteHandler).Methods(http.MethodPut)
	router.HandleFunc("/notes/{uuid}", handlers.DeleteNoteHandler).Methods(http.MethodDelete)

	// Friends endpoint
	router.HandleFunc("/friends", handlers.PutFriendHandler).Methods(http.MethodPut)
	router.HandleFunc("/friends/{ownerID}", handlers.GetAllFriendHandler).Methods(http.MethodGet)

	router.HandleFunc("/friends/search/{email}", handlers.GetOneFriendHandler).Methods(http.MethodGet)
	router.HandleFunc("/friends/{uuid}", handlers.DeleteFriendHandler).Methods(http.MethodDelete)

	log.Println("Server is running on port", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}