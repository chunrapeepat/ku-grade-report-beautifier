package model

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/api"
)

func infoSelector(body string) *api.UserInfo {
	studentNoRegex := regexp.MustCompile(`รหัสบัญชี</th><td valign='top' nowrap>b(.*?)&nbsp;`)
	nameRegex := regexp.MustCompile(`ชื่อ-นามสกุล \(ไทย\)</th><td valign='top' nowrap>(.*?)&nbsp;`)
	facultyRegex := regexp.MustCompile(`คณะ</th><td valign='top' nowrap>(.*?)&nbsp;`)
	fieldRegex := regexp.MustCompile(`สาขาวิชา</th><td valign='top' nowrap>(.*?)\s\(.+\)&nbsp;`)
	courseTypeRegex := regexp.MustCompile(`ประเภทหลักสูตร</th><td valign='top' nowrap>(.*?)&nbsp;`)
	degreeRegex := regexp.MustCompile(`ระดับปริญญา</th><td valign='top' nowrap>(.*?)&nbsp;`)

	userInfo := api.UserInfo{
		StudentNo:    studentNoRegex.FindStringSubmatch(body)[1],
		Name:         nameRegex.FindStringSubmatch(body)[1],
		Faculty:      facultyRegex.FindStringSubmatch(body)[1],
		FieldOfStudy: fieldRegex.FindStringSubmatch(body)[1],
		CourseType:   courseTypeRegex.FindStringSubmatch(body)[1],
		Degree:       degreeRegex.FindStringSubmatch(body)[1],
	}

	return &userInfo
}

// get cookie by visit website
func getCookie() (string, error) {
	client := &http.Client{}
	resp, err := client.Get("https://std.regis.ku.ac.th/_Login.php")
	if err != nil {
		return "", err
	}
	cookie := parseCookie(resp.Header)
	return cookie, nil
}

// get http header return PHPSESSID
func parseCookie(header http.Header) string {
	setCookie := header.Get("Set-Cookie")
	cookie := strings.Split(setCookie, ";")[0]
	return cookie
}
