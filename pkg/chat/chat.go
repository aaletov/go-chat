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
	log.Printf("read message: %v\n", msg)

	if err != nil {
		log.Println("Unable to read message")
		return
	}

	err = second.WriteMessage(messageType, msg)
	log.Printf("writed message: %v\n", msg)

	if err != nil {
		log.Println("Unable to write message")
		return
	}

	return
}