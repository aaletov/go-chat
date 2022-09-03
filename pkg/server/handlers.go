package server

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/aaletov/go-chat/pkg/chat"
)

var upgrader = websocket.Upgrader{}

func WebSocketHandler(w http.ResponseWriter, r *http.Request, mgr *chat.ChatManager) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %v\n", err)
		return
	}
	log.Println("Upgraded to websocket")
	defer func() {
		if err != nil {
			c.Close()
		}
	}()

	seq, err := chat.ReadChatInitSeq(c)

	if err != nil {
		log.Printf("initseq read error: %v\n", err)
		return
	}

	mgr.Add(c, seq)
	log.Printf("Added client to manager")
	return
}
