package model

import (
	"encoding/hex"
	"fmt"
)

type MgoKeyPair struct {
	Flag            byte
	MgoAddress      string
	PublicKey       []byte
	PublicKeyBase64 string
	PrivateKey      []byte
}

func (m MgoKeyPair) String() string {
	return fmt.Sprintf("Flag: %b \nMgoAddress: %s\nPublicKey: %s\nPublicKeyBase64: %s\nPrivateKey: %s\n", m.Flag, m.MgoAddress, hex.EncodeToString(m.PublicKey), m.PublicKeyBase64, hex.EncodeToString(m.PrivateKey))
}
