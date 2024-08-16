package signer

import (
	"errors"

	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model"

	"github.com/btcsuite/btcd/btcutil/bech32"
)

func DecodeMgoPrivateKey(value string) (*model.ParsedKeypair, error) {
	prefix, words, err := bech32.Decode(value)
	if err != nil {
		return nil, err
	}
	if prefix != global.MGO_PRIVATE_KEY_PREFIX {
		return nil, global.ErrPrivatekeyPrefixInvalid
	}
	extendedSecretKey, err := bech32.ConvertBits(words, 5, 8, false)
	if err != nil {
		return nil, err
	}
	secretKey := extendedSecretKey[1:]
	signatureScheme, exists := global.SIGNATURE_FLAG_TO_SCHEME[global.Scheme(extendedSecretKey[0])]
	if !exists {
		return nil, errors.New("invalid signature scheme flag")
	}

	return &model.ParsedKeypair{
		Schema:    signatureScheme,
		SecretKey: secretKey,
	}, nil
}
func EncodeMgoPrivateKey(value []byte, scheme global.Scheme) (string, error) {
	if len(value) != global.PRIVATE_KEY_SIZE {
		return "", errors.New("invalid bytes length")
	}

	privKeyBytes := append([]byte{byte(scheme)}, value...)
	words, err := bech32.ConvertBits(privKeyBytes, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode(global.MGO_PRIVATE_KEY_PREFIX, words)
}
