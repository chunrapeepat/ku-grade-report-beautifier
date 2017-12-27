package app

import "net/http"

func viewGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/view" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to view page."))
}
