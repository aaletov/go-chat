package server

import (
	"log"
	"io"
	"net/http"
	"encoding/json"
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

	_, reader, err := c.NextReader()
	buf, err := io.ReadAll(reader)

	if err != nil {
		log.Printf("read: %v\n", err)
		return
	}

	var seq chat.ChatInitSequence
	err = json.Unmarshal(buf, &seq)
	log.Println(seq)

	if err != nil {
		log.Printf("unmarshal: %v\n", err)
		return
	}

	mgr.Add(c, seq)
	log.Printf("Added client to manager")
	return
}