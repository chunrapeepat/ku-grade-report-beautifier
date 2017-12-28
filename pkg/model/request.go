package model

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/chunza2542/ku-grade-report-beautifier/pkg/api"
)

// Login login to regis ku account
// return sessionId and error
func Login(username, password, zone string) (string, error) {
	client := &http.Client{}
	// form data
	form := url.Values{}
	form.Add("form_username", username)
	form.Add("form_password", password)
	form.Add("zone", zone)
	// open new request
	req, err := http.NewRequest(
		"POST",
		"https://std.regis.ku.ac.th/_Login.php",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return "", err
	}
	// get default session & set cookie
	cookie, err := getCookie()
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	// check login from body and return sessionid
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(api.ToUTF8(body))
	if strings.Contains(bodyString, "formlogin") {
		return "", api.ErrLoginFailed
	}
	return cookie, nil
}

// GetUserInformation get all the user information except grade
// student no, name, faculty, field of study, degree
func GetUserInformation(cookie string) (*api.UserInfo, error) {
	client := &http.Client{}
	// open new request
	req, err := http.NewRequest(
		"GET",
		"https://std.regis.ku.ac.th/_Member_Information.php",
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// get information from body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(api.ToUTF8(body))
	return infoSelector(bodyString), nil
}

// Logout destroy sessionid from website
func Logout(cookie string) error {
	client := &http.Client{}
	_, err := client.Get("https://std.regis.ku.ac.th/_Logout.php")
	if err != nil {
		return err
	}
	return nil
}

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
