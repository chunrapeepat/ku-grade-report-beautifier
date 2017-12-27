package app

import "net/http"

func viewGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to view page."))
}
