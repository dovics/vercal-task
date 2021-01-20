package api

import (
	"encoding/json"
	"net/http"

	"github.com/my-Sakura/go-yuque-api/api"
)

// GetTaskInfo -
func GetTaskInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}

	values := r.URL.Query()
	Slug, ok := values["slug"]
	if !ok || len(Slug) == 0 {
		w.WriteHeader(404)
		return
	}

	doc := api.GetDocumentInfo(TOKEN, NAMESPACE, Slug[0])

	resBody, err := json.Marshal(doc.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
