package main

import (
	"log"
	"fmt"
	//"sync"
	"net/http"
	"github.com/aaletov/go-chat/api"
	"github.com/aaletov/go-chat/utils/httputil"
)

const (
	port = 8080
	maxBodySize = 1048576
	//empty void
)

func main() {
	log.Println("Starting server")

	//waitingClients := new(sync.Map)

	http.HandleFunc("/initChat", func(w http.ResponseWriter, r *http.Request) {
		status, msg := httputil.ValidateContentType(w, r, "application/json")

		if status != http.StatusOK {
			http.Error(w, msg, status)
			return
		}
		
		var initRequest api.InitChatRequest
		status, msg = httputil.Unmarshal(w, r, &initRequest)

		if status != http.StatusOK {
			http.Error(w, msg, status)
			return
		}

		log.Printf("Register request data: %v", initRequest.Key)	

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))	
	log.Println("Exiting server...")
}