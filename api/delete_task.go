package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// DeleteTask -
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	var reqBody struct {
		DocID int `json:"doc_id"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	api.DeleteDocument(TOKEN, NAMESPACE, reqBody.DocID)
	w.WriteHeader(200)
}
