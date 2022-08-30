package server

import (
	"log"
	"sync"
	"net/http"
	"github.com/aaletov/go-chat/utils/httputil"
	"github.com/aaletov/go-chat/api"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func InitChatHandler(w http.ResponseWriter, r *http.Request, waitingClients *sync.Map) {
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

	_, ok := waitingClients.Load(initRequest.RemoteKey)

	if !ok {
		waitingClients.Store(initRequest.LocalKey, initRequest.RemoteKey)
		log.Printf("The client %v haven't initialized chat; added %v to waitingClients", initRequest.RemoteKey, initRequest.LocalKey)
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("Upgraded to websocket")
	defer c.Close()
}