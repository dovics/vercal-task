package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/my-Sakura/go-yuque-api/api"
)

// GetTaskList -
func GetTaskList(w http.ResponseWriter, r *http.Request) {
	var reqBpdy = struct {
		token     string
		namespace string
	}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	docs := api.GetDocumentList(reqBody.token, reqBody.namespace)
	for _, doc := range docs.Data {
		fmt.Printf("slug: %s, title: %s\n", doc.Slug, doc.Title)
	}

	resBody, err := json.Marshal(docs.Data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
	w.Write(resBody)
}

// // GetTaskDetail -
// func GetTaskDetail(w http.ResponseWriter, r *http.Request) {
// 	var reqBpdy struct {
// 		token     string
// 		namespace string
// 	}
// 	body, _ := ioutil.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &reqBody)
// 	if err != nil {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	currentTime := time.Now().Format(time.RFC850)
// 	fmt.Fprintf(w, currentTime)
// }

// // CreateTask -
// func CreateTask(w http.ResponseWriter, r *http.Request) {
// 	var reqBpdy struct {
// 		token     string
// 		namespace string
// 	}
// 	body, _ := ioutil.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &reqBody)
// 	if err != nil {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	currentTime := time.Now().Format(time.RFC850)
// 	fmt.Fprintf(w, currentTime)
// }

// // DeleteTask -
// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	var reqBpdy struct {
// 		token     string
// 		namespace string
// 	}
// 	body, _ := ioutil.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &reqBody)
// 	if err != nil {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	currentTime := time.Now().Format(time.RFC850)
// 	fmt.Fprintf(w, currentTime)
// }

// // UpdateTask -
// func UpdateTask(w http.ResponseWriter, r *http.Request) {
// 	var reqBpdy struct {
// 		token     string
// 		namespace string
// 	}
// 	body, _ := ioutil.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &reqBody)
// 	if err != nil {
// 		w.WriteHeader(400)
// 		return
// 	}
// 	currentTime := time.Now().Format(time.RFC850)
// 	fmt.Fprintf(w, currentTime)
// }
