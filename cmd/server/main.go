package main

import (
	"log"
	"fmt"
	"sync"
	"net/http"
	"github.com/aaletov/go-chat/pkg/server"
)

const (
	port = 8080
	maxBodySize = 1048576
	//empty void
)

func main() {
	log.Println("Starting server")

	waitingClients := new(sync.Map)

	http.HandleFunc("/initChat", func(w http.ResponseWriter, r *http.Request) {
		server.InitChatHandler(w, r, waitingClients)
	})

	http.HandleFunc("/ws", server.WebSocketHandler)

	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	log.Fatal(server.ListenAndServe())	
	log.Println("Exiting server...")
}