package signer

import (
	"encoding/hex"
	"errors"
	"github.com/jarvis0919/mgo-go-sdk/utils"

	"github.com/jarvis0919/mgo-go-sdk/config"
	"github.com/jarvis0919/mgo-go-sdk/model"

	"github.com/btcsuite/btcd/btcutil/bech32"
)

func DecodeMgoPrivateKey(value string) (*model.ParsedKeypair, error) {
	prefix, words, err := bech32.Decode(value)
	if err != nil {
		return nil, err
	}
	if prefix != config.MGO_PRIVATE_KEY_PREFIX {
		return nil, config.ErrPrivatekeyPrefixInvalid
	}
	extendedSecretKey, err := bech32.ConvertBits(words, 5, 8, false)
	if err != nil {
		return nil, err
	}
	secretKey := extendedSecretKey[1:]
	signatureScheme, exists := config.SIGNATURE_FLAG_TO_SCHEME[config.Scheme(extendedSecretKey[0])]
	if !exists {
		return nil, errors.New("invalid signature scheme flag")
	}

	return &model.ParsedKeypair{
		Schema:    signatureScheme,
		SecretKey: secretKey,
	}, nil
}
func EncodeMgoPrivateKey(value []byte, scheme config.Scheme) (string, error) {
	if len(value) != config.PRIVATE_KEY_SIZE {
		return "", errors.New("invalid bytes length")
	}

	privKeyBytes := append([]byte{byte(scheme)}, value...)
	words, err := bech32.ConvertBits(privKeyBytes, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode(config.MGO_PRIVATE_KEY_PREFIX, words)
}

func PublicKeyToMgoAddressDevNet(publicKey []byte, schema config.Scheme) (string, error) {
	if schema != config.Ed25519Flag {
		return "", errors.New("invalid signature scheme flag")
	}
	tmp := append([]byte{byte(config.Ed25519Flag)}, publicKey...)
	hexHash := utils.Blake2bv1(tmp)
	return "0x" + hex.EncodeToString(hexHash)[:config.AccountAddress32Length*2], nil
}

func PublicKeyToMgoAddressTestNet(publicKey []byte, schema config.Scheme) (string, error) {
	if schema != config.Ed25519Flag {
		return "", errors.New("invalid signature scheme flag")
	}
	inputBytes := append([]byte{byte(config.Ed25519Flag)}, publicKey...)
	return "0x" + hex.EncodeToString(utils.Keccak256(inputBytes))[:config.MGO_ADDRESS_LENGTH], nil
}
