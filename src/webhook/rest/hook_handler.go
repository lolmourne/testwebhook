package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewHookHandler() http.Handler {

	mux := httprouter.New()

	mux.POST("/webhook", AddSendbirdHook())

	return mux
}
