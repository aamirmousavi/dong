package sms

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const _URL = "https://raygansms.com/SendMessageWithCode.ashx"

var (
	username = ""
	password = ""
)

func InitSetUsernameAndPassword(_username, _password string) {
	username = _username
	password = _password
}
func Send(
	mobile, message string,
) error {
	request, err := http.NewRequest("GET", _URL, nil)
	if err != nil {
		return nil
	}
	q := url.Values{}
	q.Add("UserName", username)
	q.Add("Password", password)
	q.Add("Mobile", mobile)
	q.Add("Message", message)
	request.URL.RawQuery = q.Encode()
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	reponseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("status wasn't 200 body: %v", string(reponseBody))
	}
	return nil
}
