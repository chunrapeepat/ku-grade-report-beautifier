package app

import (
	"net/http"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/api"
	"github.com/chunza2542/ku-grade-report-beautifier/pkg/model"
	"github.com/chunza2542/ku-grade-report-beautifier/pkg/view"
)

func indexGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		view.Index(w, r, &view.IndexData{})
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")
		// Login to endpoint using Nontri account
		session, err := model.Login(username, password, "0")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Get User information & Grade
		userInfo, err := model.GetUserInformation(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		grade, err := model.GetGradeReport(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Logout & Destroy Session
		err = model.Logout(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// calculate gpa
		gpa, err := api.CalculateGPA(grade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Parse Data
		view.Report(w, r, &view.ReportData{
			UserInfo:   userInfo,
			CourseInfo: grade,
			GPA:        gpa,
		})
	}
}
