package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/chat/text_completion", handleSingleCompletion).Methods("POST")
	r.HandleFunc("/chat/chat_completion_stateless", handleChatChatCompletionStateless).Methods("POST")
	r.HandleFunc("/chat/chat_completion", handleChatChatCompletion).Methods("POST")
	r.HandleFunc("/chat/create_chat_id", createChatId).Methods("POST")
	r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
}
