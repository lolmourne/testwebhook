package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {

	log.SetFlags(log.Llongfile | log.Ldate)

	mux := newRouter()

	http.ListenAndServe(":8080", mux)
}

func newRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/webhook", webhookHandler).Methods("POST")
	return r
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var sendBird SendBird

	err := decoder.Decode(&sendBird)

	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	log.Println(sendBird.AppID)

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
