package keyutil

import (
	"log"
	"fmt"
	"os"
	"errors"
	"math/big"
	"crypto/ecdsa"
	"crypto/elliptic"
)

func MarshalECDSAPrivate(curve elliptic.Curve, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	publicKey := privateKey.PublicKey
	marshalledKey := elliptic.Marshal(curve, publicKey.X, publicKey.Y)
	keySize := len(marshalledKey)

	if keySize > 255 {
		err := errors.New("can't write keys larger than 255 bytes")
		log.Println(err)
		return nil, err
	}

	byteKeySize := byte(keySize)
	keyWithSize := append([]byte{byteKeySize}, marshalledKey...)
	marshalledD, err := privateKey.D.MarshalText()

	if err != nil {
		log.Printf("unable to marshal private key: %v\n", err)
		return nil, err
	}

	dSize := len(marshalledD)
	if dSize > 255 {
		err := errors.New("can't write D larger than 255 bytes")
		log.Println(err)
		return nil, err
	}

	byteDSize := byte(dSize)
	dWithSize := append([]byte{byteDSize}, marshalledD...)

	return append(keyWithSize, dWithSize...), nil
}

// read reads message from file which is preceded with size
func read(f *os.File) ([]byte, error) {
	sizeBuf := make([]byte, 1)
	_, err := f.Read(sizeBuf)

	if err != nil {
		log.Printf("read error: %v", err)
		return nil, err
	}

	valBuf := make([]byte, int(sizeBuf[0]))
	_, err = f.Read(valBuf)

	if err != nil {
		log.Printf("read error: %v", err)
		return nil, err
	}

	return valBuf, nil
}

func ReadKey(keyName string) (keyBuf, dBuf []byte, err error) {
	f, err := os.OpenFile(fmt.Sprintf("%v.dat", keyName), os.O_RDONLY, 0600)

	if err != nil {
		log.Printf("unable to open file: %v", err)
		return nil, nil, err
	}
	defer f.Close()

	keyBuf, _ = read(f)
	dBuf, _ = read(f)

	return
}

func UnmarshalECDSAPrivate(curve elliptic.Curve, keyBuf, dBuf []byte) (*ecdsa.PrivateKey, error) {
	x, y := elliptic.Unmarshal(curve, keyBuf)
	var z *big.Int = new(big.Int)
	err := z.UnmarshalText(dBuf)

	if err != nil {
		log.Println("unable to unmarshal key: %v", err)
		return nil, err
	}

	publicKey := ecdsa.PublicKey{curve, x, y}

	return &ecdsa.PrivateKey{publicKey, z}, nil
}