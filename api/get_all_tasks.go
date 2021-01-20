package api

import (
	"encoding/json"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// GetTaskList -
func GetTaskList(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}

	docs := api.GetDocumentList(TOKEN, NAMESPACE)

	resBody, err := json.Marshal(docs.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
