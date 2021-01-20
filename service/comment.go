package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var template = `{"commentable_type":"Doc","commentable_id":%d,"body_asl":"<u0021doctype lake><meta name=\"doc-version\" content=\"%s\" /><meta name=\"viewport\" content=\"fixed\" /><meta name=\"typography\" content=\"traditional\" /><p data-lake-id=\"a20d224c568e48b9d67847a2c66a8c01_p_0\">1</p><u0021bd15a6bae5411959542913d61de0a2e6d1563e4abb7b11199f92a57e66523fc8>","format":"lake"}`

// GetComments -
func GetComments(session string, docID int) (*ResponseCommentInfoSerializer, error) {
	url := fmt.Sprintf("https://www.yuque.com/api/comments?commentable_id=%d&commentable_type=Doc&include_reactions=true&include_to_user=true", docID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("cookie", fmt.Sprintf("_yuque_session=%s", session))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := ResponseCommentInfoSerializer{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddComment -
func AddComment(session string, docID int, content string) (*ResponseAddCommentSerializer, error) {
	url := "https://www.yuque.com/api/comments"
	method := "POST"

	payloadString := fmt.Sprintf(template, docID, content)
	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Add("cookie", fmt.Sprintf("_yuque_session=%s", session))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := ResponseAddCommentSerializer{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
