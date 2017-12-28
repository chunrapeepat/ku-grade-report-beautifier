package view

import "net/http"

type IndexData struct {
}

type ReportData struct {
}

func Index(w http.ResponseWriter, r *http.Request, data *IndexData) {
	render(tpIndex, w, data)
}

func Report(w http.ResponseWriter, r *http.Request, data *ReportData) {
	render(tpReport, w, data)
}
