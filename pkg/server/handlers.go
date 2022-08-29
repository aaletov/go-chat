package server

import (
	"log"
	"net/http"
	"github.com/aaletov/go-chat/utils/httputil"
	"github.com/aaletov/go-chat/api"
)

func InitChatHandler(w http.ResponseWriter, r *http.Request) {
	status, msg := httputil.ValidateContentType(w, r, "application/json")

	if status != http.StatusOK {
		http.Error(w, msg, status)
		log.Printf("Content-Type is invalid: %v", msg)
		return
	}
	
	var initRequest api.InitChatRequest
	status, msg = httputil.Unmarshal(w, r, &initRequest)

	if status != http.StatusOK {
		http.Error(w, msg, status)
		log.Printf("Unable to unmarshal data: %v", msg)
		return
	}

	log.Printf("Initialized waiting chat for key: %v %v", initRequest.Key.N, initRequest.Key.E)	

	w.WriteHeader(http.StatusOK)
}