package api

import (
	"crypto/rsa"
)

type InitChatRequest struct {
	Key rsa.PublicKey `json:"key"`
}