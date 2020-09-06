package main

import (
	"github.com/gorilla/mux"
	"gratitude/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")

	router.HandleFunc("/notes", handlers.GetAllNoteHandler).Methods(http.MethodGet)
	router.HandleFunc("/notes/{uuid}", handlers.GetOneNoteHandler).Methods(http.MethodGet)
	router.HandleFunc("/notes/{uuid}", handlers.PostNoteHandler).Methods(http.MethodPost)
	router.HandleFunc("/notes/{uuid}", handlers.PutNoteHandler).Methods(http.MethodPut)
	router.HandleFunc("/notes/{uuid}", handlers.DeleteNoteHandler).Methods(http.MethodDelete)


	router.HandleFunc("/friends", handlers.GetAllFriendHandler).Methods(http.MethodGet)
	router.HandleFunc("/friends/{uuid}", handlers.GetOneNoteHandler).Methods(http.MethodGet)
	router.HandleFunc("/friends/{uuid}", handlers.PostFriendHandler).Methods(http.MethodPost)
	router.HandleFunc("/friends/{uuid}", handlers.PutFriendHandler).Methods(http.MethodPut)
	router.HandleFunc("/friends/{uuid}", handlers.DeleteFriendHandler).Methods(http.MethodDelete)

	log.Println("Server is running on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		panic(err)
	}
}