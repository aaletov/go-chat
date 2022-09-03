package chat

import (
	"io"
	"encoding/json"
	"log"
	"crypto/elliptic"
	"crypto/ecdsa"
	"math/big"
	"github.com/gorilla/websocket"
)

// Transfers points on eliptic curve P224
type ChatInitSequence struct {
	LocalKey []byte `json:"localKey"`
	RemoteKey []byte `json:remoteKey`
}

func ReadChatInitSeq(c *websocket.Conn) (*ChatInitSequence, error) {
	_, reader, err := c.NextReader()
	buf, err := io.ReadAll(reader)

	if err != nil {
		log.Printf("read: %v\n", err)
		return nil, err
	}

	var seq ChatInitSequence
	err = json.Unmarshal(buf, &seq)
	log.Println(seq)

	if err != nil {
		log.Printf("unmarshal: %v\n", err)
		return nil, err
	}

	return &seq, nil
}

type ChatKeyPair struct {
	LocalKey ecdsa.PublicKey
	RemoteKey ecdsa.PublicKey
}

var curve = elliptic.P224()

func NewKeyPair(seq ChatInitSequence) ChatKeyPair {
	lX, lY := elliptic.Unmarshal(curve, seq.LocalKey)
	rX, rY := elliptic.Unmarshal(curve, seq.RemoteKey)
	return ChatKeyPair{
		LocalKey: ecdsa.PublicKey{curve, lX, lY},
		RemoteKey: ecdsa.PublicKey{curve, rX, rY},
	}
}

func (c ChatKeyPair) GetChatIdentifier() *big.Int {
	log.Println(c)
	localNum := appendBigInt(c.LocalKey.X, c.LocalKey.Y)
	remoteNum := appendBigInt(c.RemoteKey.X, c.RemoteKey.Y)
	
	chatId := ellipticPairingFunction(localNum, remoteNum)
	log.Printf("chat id: %v\n", chatId)
	return chatId
}

func appendBigInt(l, r *big.Int) *big.Int {
	num := new(big.Int)
	num.SetBytes(l.Bytes())
	buf := num.Append(r.Bytes(), 10)
	num.SetBytes(buf)
	return num
}

func ellipticPairingFunction(x, y *big.Int) *big.Int {
	// x^{2}+y^{2}+xy+x+y
	xPower := new(big.Int)
	xPower.Mul(x, x)
	yPower := new(big.Int)
	yPower.Mul(y, y)
	xyMul := new(big.Int)
	xyMul.Mul(x, y)
	powerSum := new(big.Int)
	powerSum.Add(xPower, yPower)
	firstPowSum := new(big.Int)
	firstPowSum.Add(xyMul, x)
	firstPowSum.Add(firstPowSum, y)
	res := new(big.Int)
	res.Add(powerSum, firstPowSum)

	return res 
}