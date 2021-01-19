package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// GetTaskDetail -
func GetTaskDetail(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Token     string `json:"token"`
		Namespace string `json:"namespace"`
		Slug      string `json:"slug"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	doc := api.GetDocumentInfo(reqBody.Token, reqBody.Namespace, reqBody.Slug)

	resBody, err := json.Marshal(doc.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
