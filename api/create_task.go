package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// CreateTask -
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Token     string `json:"token"`
		Namespace string `json:"namespace"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	doc := api.CreateDocument(reqBody.Token, reqBody.Namespace, reqBody.Slug, reqBody.Title)

	resBody, err := json.Marshal(doc.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
