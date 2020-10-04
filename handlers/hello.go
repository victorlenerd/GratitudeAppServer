package handlers

import "net/http"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	body := "Gratitude App Server Running"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}