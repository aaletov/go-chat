package chat

import (
	"crypto/ecdsa"
	"github.com/gorilla/websocket"
)

type ClientConnection struct {
	LocalKey *ecdsa.PublicKey
	Connection *websocket.Conn
}