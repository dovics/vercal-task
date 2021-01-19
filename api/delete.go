package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// DeleteTask -
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var reqBpdy struct {
		token     string
		namespace string
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}
