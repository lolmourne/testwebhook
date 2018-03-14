package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {

	log.SetFlags(log.Llongfile | log.Ldate)

	mux := newRouter()

	http.ListenAndServe(":8756", mux)
}

func newRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/webhook", webhookHandler).Methods("POST")
	return r
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	var request []string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	request = append(request, fmt.Sprintf("Host: %v", r.Host))

	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	decoder := json.NewDecoder(r.Body)

	var sendBird SendBird

	err := decoder.Decode(&sendBird)

	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	log.Println(sendBird.AppID)
	log.Println(request)

}

type SendBird struct {
	AppID    string `json:"app_id"`
	Category string `json:"category"`
	Type     string `json:"type"`
	SDK      string `json:"sdk"`
	Sender   Sender `json:"sender"`
}

type Sender struct {
	UserID     int64  `json:"user_id"`
	Nickname   string `json:"nickname"`
	ProfileURL string `json:"profile_url"`
}
