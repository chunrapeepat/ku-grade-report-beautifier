package model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

const OFFSET = 0xd60
const WIDTH = 3

func ToUTF8(tis620bytes []byte) []byte {
	l := findOutputLength(tis620bytes)
	output := make([]byte, l)

	index := 0
	buffer := make([]byte, WIDTH)
	for _, c := range tis620bytes {
		if !isThaiChar(c) {
			output[index] = c

			index++
		} else {
			utf8.EncodeRune(buffer, int32(c)+OFFSET)
			output[index] = buffer[0]
			output[index+1] = buffer[1]
			output[index+2] = buffer[2]

			index += 3
		}
	}
	return output
}

func findOutputLength(tis620bytes []byte) int {
	outputLen := 0
	for i, _ := range tis620bytes {
		if isThaiChar(tis620bytes[i]) {
			outputLen += WIDTH //always 3 bytes for thai char
		} else {
			outputLen += 1
		}
	}
	return outputLen
}

func isThaiChar(c byte) bool {
	return (c >= 0xA1 && c <= 0xDA) || (c >= 0xDF && c <= 0xFB)
}

// Login login to regis ku account
// return sessionId and error
func Login(username, password, zone string) (*string, error) {
	form := url.Values{}
	form.Add("form_username", username)
	form.Add("form_password", password)
	form.Add("zone", zone)
	resp, err := http.PostForm(
		"https://std.regis.ku.ac.th/_Login.php",
		form,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(ToUTF8(body))
	if strings.Contains(bodyString, "formlogin") {
		fmt.Println("Login failed")
	}
	fmt.Println("Success~", resp.Header.Get("Set-Cookie"))
	return nil, nil
}
