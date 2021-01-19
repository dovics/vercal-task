package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// GetTaskList -
func GetTaskList(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Token     string `json:"token"`
		Namespace string `json:"namespace"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	docs := api.GetDocumentList(reqBody.Token, reqBody.Namespace)

	resBody, err := json.Marshal(docs.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
