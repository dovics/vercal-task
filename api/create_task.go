package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/my-Sakura/go-yuque-api/api"
)

// TOKEN -
var TOKEN string

// NAMESPACE -
var NAMESPACE string

func init() {
	TOKEN = os.Getenv("TOKEN")
	NAMESPACE = os.Getenv("NAMESPACE")
}

// CreateTask -
func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	var reqBody struct {
		Title string `json:"title"`
		Slug  string `json:"slug"`
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	doc := api.CreateDocument(TOKEN, NAMESPACE, reqBody.Slug, reqBody.Title)

	resBody, err := json.Marshal(doc.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}
