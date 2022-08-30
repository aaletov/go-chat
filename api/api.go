package api

import (
	"crypto/rsa"
)

type InitChatRequest struct {
	LocalKey rsa.PublicKey `json:"localKey"`
	RemoteKey rsa.PublicKey `json:remoteKey`
}