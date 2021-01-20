package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dovics/task-manager/service"
)

// GetTaskContent -
func GetTaskContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(404)
		return
	}

	values := r.URL.Query()
	docIDStrings, ok := values["doc_id"]
	if !ok || len(docIDStrings) == 0 {
		w.WriteHeader(404)
		return
	}

	docID, err := strconv.Atoi(docIDStrings[0])
	if err != nil {
		w.WriteHeader(404)
		return
	}

	comments, err := service.GetComments(SESSION, docID)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	resBody, err := json.Marshal(comments.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
