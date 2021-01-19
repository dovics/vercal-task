package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// DeleteTask -
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Token     string `json:"token"`
		Namespace string `json:"namespace"`
		DocID     int    `json:"doc_id"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	api.DeleteDocument(reqBody.Token, reqBody.Namespace, reqBody.DocID)
	w.WriteHeader(200)
}
