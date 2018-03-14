package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AddSendbirdHook() httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	}

}
