package main

import (
	"net/http"
	"log"
	//"fmt"
	"crypto/rsa"
	"crypto/rand"
	"flag"
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/aaletov/go-chat/api"
)

const (
	lport = 8082
	rport = 8081
	remoteURL = "http://localhost:8080/initChat"
	wsEndpoint = "ws://localhost:8080/ws"
	keySize = 256
)

var (
	msg = flag.String("msg", "", "")
)

func main() {
	log.Println("Starting client")

	randReader := rand.Reader
	privateKey, err := rsa.GenerateKey(randReader, keySize)

	if err != nil {
		panic(err)
	}

	lKey := privateKey.Public().(*rsa.PublicKey)
	rKey := privateKey.Public().(*rsa.PublicKey)
	initRequest := api.InitChatRequest{*lKey, *rKey}
	requestJSON, err := json.Marshal(initRequest)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post(remoteURL, "application/json", bytes.NewBuffer(requestJSON))

	c, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Fatal("dial:", err)
	}
	log.Println("Opened websocket")
	defer c.Close()

	if err != nil {
		panic(err)
	}
	log.Println(resp)

	for {
		err = c.WriteMessage(websocket.TextMessage, []byte(*msg))
		if err != nil {
			log.Println(err)
			return
		}
	}
	//log.Printf("Public key is: %v", privateKey.Public())

	// laddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", lport))
	// raddr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", rport))
	// conn, _ := net.DialTCP("tcp", laddr, raddr)

	// log.Printf("local addr: %v, remote addr: %v\n", conn.LocalAddr(), conn.RemoteAddr())
	log.Println("Exiting client..")
}