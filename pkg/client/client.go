package client

import (
	"crypto/ecdsa"
)

type Client struct {
	Key *ecdsa.PrivateKey
}