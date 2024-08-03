package model

import (
	"encoding/hex"
	"fmt"
)

type ParsedKeypair struct {
	Schema    string
	SecretKey []byte
}

func (p ParsedKeypair) String() string {
	return fmt.Sprintf("Schema: %s \nSecretKey: %s", p.Schema, hex.EncodeToString(p.SecretKey))
}
