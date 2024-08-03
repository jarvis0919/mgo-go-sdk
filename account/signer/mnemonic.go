package signer

import (
	"errors"
	"fmt"
	"mgo-go-sdk/global"

	"github.com/tyler-smith/go-bip39"
)

func NewSignertWithMnemonicTestNet(mnemonic string, keytype global.Keytype) (Signer, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	derivation, ok := global.DERIVATION_PATH_TESTNET[keytype]
	if !ok {
		return nil, errors.New("invalid signature scheme flag")
	}
	key, err := DeriveForPath(derivation, seed)
	if err != nil {
		return nil, err
	}
	signer, err := newEd25519SignerFromSeed(key.Key)
	return signer, err
}
func NewSignertWithMnemonicDevNet(mnemonic string, keytype global.Keytype) (Signer, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	fmt.Println(bip39.NewMnemonic(seed))
	if err != nil {
		return nil, err
	}

	derivation, ok := global.DERIVATION_PATH_DEVNET[keytype]
	if !ok {
		return nil, errors.New("invalid signature scheme flag")
	}
	key, err := DeriveForPath(derivation, seed)
	if err != nil {
		return nil, err
	}
	signer, err := newEd25519SignerFromSeed(key.Key)
	return signer, err
}
