package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/aaletov/go-chat/pkg/server"
	"github.com/aaletov/go-chat/pkg/chat"
)

const (
	port = 8080
	maxBodySize = 1048576
	//empty void
)

func main() {
	log.Println("Starting server")

	mgr := chat.NewManager()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.WebSocketHandler(w, r, mgr)
	})

	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	go mgr.Start()

	log.Fatal(server.ListenAndServe())	
	log.Println("Exiting server...")
}