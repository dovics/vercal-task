package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dovics/task-manager/service"
)

// UpdateTaskContent -
func UpdateTaskContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	var reqBody struct {
		DocID   int    `json:"doc_id"`
		Content string `json:"content"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	comment, err := service.AddComment(SESSION, reqBody.DocID, reqBody.Content)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	resBody, err := json.Marshal(comment.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
