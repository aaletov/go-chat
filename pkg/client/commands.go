package client

import (
	"log"
	"fmt"
	"os"
	"time"
	"errors"
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"encoding/base64"
	"github.com/gorilla/websocket"
	"github.com/aaletov/go-chat/utils/keyutil"
	"github.com/aaletov/go-chat/pkg/chat"
)

const (
	wsEndpoint = "ws://localhost:8080/ws"
)

var (
	curve = elliptic.P224()
	encoding = base64.StdEncoding
)

func (c *Client) GenKey() error {
	var keyName string
	_, err := fmt.Scan(&keyName)

	if err != nil {
		return err
	}

	randReader := rand.Reader
	privateKey, err := ecdsa.GenerateKey(curve, randReader)

	if err != nil {
		log.Printf("unable to generate private key: %v\n", err)
		return err
	}

	marshalledKey, _ := keyutil.MarshalECDSAPrivate(curve, privateKey)
	f, err := os.OpenFile(fmt.Sprintf("%v.dat", keyName), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Printf("unable to open file: %v\n", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(marshalledKey)
	if err != nil {
		log.Printf("unable to write key to file: %v\n", err)
		return err
	}

	log.Println("Writted key")

	return nil
}

func (c *Client) UseKey() error {
	var keyName string
	_, err := fmt.Scan(&keyName)

	if err != nil {
		return err
	}

	keyBuf, dBuf, err := keyutil.ReadKey(keyName)

	if err != nil {
		log.Printf("unable to read key %v: %v", keyName, err)
		return err
	}

	key, err := keyutil.UnmarshalECDSAPrivate(curve, keyBuf, dBuf)

	if err != nil {
		log.Printf("unable to unmarshal key: %v", err)
		return err
	}

	c.Key = key

	return nil
}

func (c *Client) GetKey() error {
	if c.Key == nil {
		return errors.New("No key to use")
	}

	marshalledKey := elliptic.Marshal(curve, c.Key.X, c.Key.Y)
	base64Key := make([]byte, encoding.EncodedLen(len(marshalledKey)))
	encoding.Encode(base64Key, marshalledKey)
	fmt.Println(string(base64Key))

	return nil
}

func (c *Client) Chat() error {
	var remoteKey string
	_, err := fmt.Scan(&remoteKey)

	if err != nil {
		return err
	}

	byteKey := []byte(remoteKey)
	marshalledKey := make([]byte, encoding.DecodedLen(len(byteKey)))
	encoding.Decode(marshalledKey, byteKey)

	initRequest := chat.ChatInitSequence{
		elliptic.Marshal(curve, c.Key.X, c.Key.Y),
		marshalledKey,
	}
	requestJSON, err := json.Marshal(initRequest)

	if err != nil {
		log.Printf("json marshal:", err)
		return err
	}

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Printf("dial:", err)
		return err
	}
	log.Println("Opened websocket")
	defer func() {
		deadline := time.Now().Add(time.Duration(1000000))
		conn.WriteControl(websocket.CloseMessage, []byte{}, deadline)
		conn.Close()
	}()

	err = conn.WriteMessage(websocket.TextMessage, []byte(requestJSON))

	if err != nil {
		log.Printf("write message:", err)
		return err
	}

	close := make(chan bool)
	go func() {
		var msg string
		for {
			_, err = fmt.Scanln(&msg)

			if err != nil {
				fmt.Println("interrupted")
				break
			}

			err = conn.WriteMessage(websocket.TextMessage, []byte(msg))

			if err != nil {
				log.Println(err)
				break
			}
		}

		close <- true
	}()

	go func () {
		for {
			_, byteMsg, err := conn.ReadMessage()

			if websocket.IsCloseError(err) {
				fmt.Println("Connection closed, message was not delivered")
				break
			}

			if err != nil {
				log.Println(err)
				break
			}
	
			fmt.Printf("incoming: %v\n", string(byteMsg))
		}

		close <- true
	}()

	<- close

	fmt.Println("Chat closed")

	return nil
}