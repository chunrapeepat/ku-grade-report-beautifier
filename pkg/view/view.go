package view

import (
	"net/http"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/api"
)

// IndexData view
type IndexData struct {
}

// ReportData view
type ReportData struct {
	UserInfo   *api.UserInfo
	CourseInfo []*api.CourseInfo
	GPA        float32
}

// Index render index template to view
func Index(w http.ResponseWriter, r *http.Request, data *IndexData) {
	render(tpIndex, w, data)
}

// Report render report template to view
func Report(w http.ResponseWriter, r *http.Request, data *ReportData) {
	render(tpReport, w, data)
}
