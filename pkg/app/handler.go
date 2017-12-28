package app

import (
	"net/http"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/view"
)

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	view.Index(w, r, &view.IndexData{})
}
