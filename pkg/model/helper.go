package model

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/api"
)

// get all user information using regex
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

// select all grade list using regex
func gradeSelector(body string) []*api.CourseInfo {
	tableRegex := regexp.MustCompile(`<td valign=top  align=left  width=325px> <table border=0 class=table  width=100% cellspacing=0 cellpadding=0><tr bgcolor=#006633><th align=center class=th_cs width=25%>Course<br> Code</th><th align=center class=th_cs  width=40% >Course Title</th><th align=center class=th_cs  width=20%>Grade</th><th align=center class=th_cs  width=15%>Credit</th></tr><tr><td  colspan=4 class=head_sm align=center>First Semester 2017</td></tr>(.*?)<tr><td align=left colspan=4 class=gpa_sm><dd><dd>`)
	table := tableRegex.FindStringSubmatch(body)[1]

	var report []*api.CourseInfo

	trRegex := regexp.MustCompile(`<tr>(.*?)</tr>`)
	for _, tr := range trRegex.FindAllString(table, -1) {
		var courseInfo api.CourseInfo
		tdRegex := regexp.MustCompile(`<td (nowrap|align=center)? class=tr>(<FONT COLOR="#(.+)">)?(.*?)(</FONT>)?</td>`)
		td := tdRegex.FindAllStringSubmatch(tr, -1)
		courseInfo.Code = td[0][4]
		courseInfo.Title = td[1][4]
		courseInfo.Grade = td[2][4]
		courseInfo.Credit = td[3][4]
		report = append(report, &courseInfo)
	}

	return report
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
