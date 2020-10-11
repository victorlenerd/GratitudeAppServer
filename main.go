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

	router.Use(mux.CORSMethodMiddleware(router))

	router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Content-Type", "application/json")
			handler.ServeHTTP(w, r)
		})
	})
	
	// Feeds endpoint
	router.HandleFunc("/feeds/{uuid}", handlers.FeedsHandler).Methods(http.MethodGet)

	// Friends endpoint
	router.HandleFunc("/friends", handlers.PutFriendHandler).Methods(http.MethodPut)
	router.HandleFunc("/friends/{ownerID}", handlers.GetAllFriendHandler).Methods(http.MethodGet)
	router.HandleFunc("/friends/{uuid}", handlers.DeleteFriendHandler).Methods(http.MethodDelete)

	// Notes endpoint
	router.HandleFunc("/notes", handlers.GetAllNoteHandler).Methods(http.MethodGet)
	router.HandleFunc("/notes", handlers.PutNoteHandler).Methods(http.MethodPut)
	router.HandleFunc("/notes/{uuid}", handlers.DeleteNoteHandler).Methods(http.MethodDelete)

	// User Firebase Messaging Tokens Endpoint
	router.HandleFunc("/search", handlers.PutUserToken).Methods(http.MethodPut)

	// Search Endpoint
	router.HandleFunc("/search", handlers.SearchFriendHandler).Methods(http.MethodGet)

	// Welcome
	router.HandleFunc("/", handlers.HelloHandler).Methods(http.MethodGet)

	log.Println("Server is running on port", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}