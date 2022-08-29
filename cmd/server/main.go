package main

import (
	"log"
	"fmt"
	//"sync"
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

	//waitingClients := new(sync.Map)

	http.HandleFunc("/initChat", server.InitChatHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))	
	log.Println("Exiting server...")
}