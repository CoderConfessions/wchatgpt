package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/chat/text-completion", HandleSingleCompletion).Methods("POST")
	r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
}
