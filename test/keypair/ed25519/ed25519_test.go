package main

import (
	"testing"

	"github.com/jarvis0919/mgo-go-sdk/account/signer/ed25519"
)

func TestEd25519(t *testing.T) {

	sig, err := ed25519.NewEd25519Signer()
	if err != nil {
		panic(err)
	}
	t.Log(sig.PrivateKeyHex())
	t.Log(sig.PublicKeyHex())
	t.Log(sig.MgoPrivateKey())

	sig, err = ed25519.NewEd25519SignerFromMgoPrivatekey(sig.MgoPrivateKey())
	if err != nil {
		panic(err)
	}
	t.Log(sig.PrivateKeyHex())
	t.Log(sig.PublicKeyHex())
	t.Log(sig.MgoPrivateKey())

	sig, err = ed25519.NewEd25519SignerFromPrivateKey(sig.PrivateKeyHex())
	if err != nil {
		panic(err)
	}
	t.Log(sig.PrivateKeyHex())
	t.Log(sig.PublicKeyHex())
	t.Log(sig.MgoPrivateKey())

	data := sig.Sign([]byte("hello world"))

	t.Log(data)
	t.Log(ed25519.Verify(sig.PublicKeyBytes(), []byte("hello world"), data))
}
