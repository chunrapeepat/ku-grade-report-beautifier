package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
func GetUserInformation(cookie string) {
	client := &http.Client{}
	// open new request
	req, err := http.NewRequest(
		"GET",
		"https://std.regis.ku.ac.th/_Member_Information.php",
		nil,
	)
	if err != nil {
		return
	}
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// get information from body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	bodyString := string(api.ToUTF8(body))
	fmt.Println(bodyString)
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
