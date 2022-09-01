package main

import (
	"log"
	"crypto/ecdsa"
	"crypto/elliptic"
	//"crypto/rand"
	"math/big"
	"flag"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/aaletov/go-chat/pkg/chat"
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

const (
	X = "22932723935884356325978933133478612719870509751359180570662911120639"
	Y = "10362621957340144616877306252686708115149192148772687366931180099725"
)

func main() {
	flag.Parse()
	log.Println("Starting client")

	//randReader := rand.Reader
	curve := elliptic.P224()
	log.Println(curve.Params())
	// privateKey, err := ecdsa.GenerateKey(curve, randReader)

	// if err != nil {
	// 	panic(err)
	// }

	// lKey := privateKey.Public().(*ecdsa.PublicKey)
	// rKey := privateKey.Public().(*ecdsa.PublicKey)

	// log.Println(lKey.X)
	// log.Println(rKey.Y)

	xBig := new(big.Int)
	xBig.SetString(X, 10)
	yBig := new(big.Int)
	yBig.SetString(Y, 10)

	log.Println(xBig)
	log.Println(yBig)
	lKey := &ecdsa.PublicKey{curve, xBig, yBig}
	rKey := &ecdsa.PublicKey{curve, xBig, yBig}

	initRequest := chat.ChatInitSequence{
		elliptic.Marshal(curve, lKey.X, lKey.Y),
		elliptic.Marshal(curve, rKey.X, rKey.Y),
	}
	requestJSON, err := json.Marshal(initRequest)

	if err != nil {
		panic(err)
	}

	c, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Fatal("dial:", err)
	}
	log.Println("Opened websocket")
	defer c.Close()

	log.Println(requestJSON)
	err = c.WriteMessage(websocket.TextMessage, []byte(requestJSON))

	if err != nil {
		panic(err)
	}

	log.Println(*msg)
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