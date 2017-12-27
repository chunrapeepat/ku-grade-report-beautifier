package app

import "net/http"

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to index page"))
}
