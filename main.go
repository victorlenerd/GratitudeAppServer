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

	ctx := context.Background()
	// Initialize Data store Client
	db.Init(ctx)

	// Feeds endpoint

	// Notes endpoint

	router.HandleFunc("/notes", handlers.GetAllNoteHandler).Methods(http.MethodGet)
	router.HandleFunc("/notes", handlers.PutNoteHandler).Methods(http.MethodPut)
	router.HandleFunc("/notes/{uuid}", handlers.DeleteNoteHandler).Methods(http.MethodDelete)

	// Friends endpoint

	router.HandleFunc("/friends", handlers.GetAllFriendHandler).Methods(http.MethodGet)
	router.HandleFunc("/friends/search/{email}", handlers.GetOneFriendHandler).Methods(http.MethodGet)

	router.HandleFunc("/friends/{uuid}", handlers.PutFriendHandler).Methods(http.MethodPut)
	router.HandleFunc("/friends/{uuid}", handlers.DeleteFriendHandler).Methods(http.MethodDelete)

	log.Println("Server is running on port", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}