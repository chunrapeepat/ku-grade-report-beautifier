package app

import "net/http"

// Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", indexGetHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}
