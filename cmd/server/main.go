package main

import (
	"log"
	"fmt"
	//"sync"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/aaletov/go-chat/api"
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
		if r.Header.Get("Content-Type") == "" {
			msg := "Content-Type is empty"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
		value := r.Header.Get("Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
		
		reader := http.MaxBytesReader(w, r.Body, maxBodySize)
		body, err := ioutil.ReadAll(reader)

		if err != nil {
			msg := "Unable to read body"
			http.Error(w, msg, http.StatusUnprocessableEntity)
			return
		}

		var registerRequest api.InitChatRequest
		err = json.Unmarshal(body, &registerRequest)

		if err != nil {
			msg := "Invalid request body"
			http.Error(w, msg, http.StatusUnprocessableEntity)
			return
		}

		log.Printf("Register request data: %v", registerRequest.Key)	

		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))	
	log.Println("Exiting server...")
}