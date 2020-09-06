package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"scheduler0/server/src/utils"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")


	log.Println("Server is running on port", port)
	err := http.ListenAndServe(utils.GetPort(), router)
	if err != nil {
		panic(err)
	}
}