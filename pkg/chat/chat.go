package chat

import (
	"log"
	"github.com/gorilla/websocket"
)

type Chat struct {
	First *websocket.Conn
	Second *websocket.Conn
}

func sendMessage(first *websocket.Conn, second *websocket.Conn) (err error) {
	messageType, msg, err := first.ReadMessage()

	if err != nil {
		log.Println("Unable to read message")
		return
	}

	err = second.WriteMessage(messageType, msg)

	if err != nil {
		log.Println("Unable to write message")
		return
	}

	return
}