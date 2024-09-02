package main

import (
	"context"
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"testing"

	"github.com/jarvis0919/mgo-go-sdk/account/keypair"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/config"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/utils"
)

var ctx = context.Background()
var devCli = client.NewMgoClient(config.MgoDevnet)

func TestSignPersonalMessage(t *testing.T) {

	sig, err := keypair.New(keypair.Options{Scheme: config.Ed25519Flag})

	if err != nil {
		panic(err)
	}
	signData := sig.SignPersonalMessage([]byte("hello world"), config.MgoDevnet)

	t.Log(signData)
	t.Log(utils.ByteArrayToBase64String(signData))

	result := keypair.VerifyPersonalMessage([]byte("hello world"), signData, config.MgoDevnet)
	t.Log(result)

	scheme := keypair.GetSignatureScheme(signData)
	signatureInfo := keypair.ParseSignatureInfo(signData, scheme)

	addressDevNet, err := signer.PublicKeyToMgoAddressDevNet(signatureInfo.PublicKey, config.Ed25519Flag)

	if sig.ToMgoAddressDevNet() != addressDevNet {
		t.Fatal("addressDevNet not match")
	}
}

func TestSignTransactionBlock(t *testing.T) {

	sig, err := keypair.New(keypair.Options{Scheme: config.Ed25519Flag})
	if err != nil {
		panic(err)
	}
	pay, err := devCli.Pay(ctx, request.PayRequest{
		Signer:      "0x19e2112d5589580c1b9ca0b682f481935f1d7d84948c0f220b8633a0e6ef9712",
		MgoObjectId: []string{"0x03e26828380b193c61554e405b2ab91e5e0db118907259bab9897ade63540a8b"},
		Recipient:   []string{"0x19e2112d5589580c1b9ca0b682f481935f1d7d84948c0f220b8633a0e6ef9712"},
		Amount:      []string{"1000"},
		Gas:         "0xfab148b55f87b19577512de8f2a2a461088ad0b14e85e28f6a63a3df90812896",
		GasBudget:   "10000000",
	})
	if err != nil {
		t.Fatal(err)
	}

	utils.JsonPrint(pay)
	data := sig.SignTransactionBlock(&pay, config.MgoDevnet)

	t.Log(data.TxBytes)
	t.Log(len(utils.DecodeBase64(data.Signature)))
	t.Log(keypair.VerifyTransactionBlock(utils.DecodeBase64(data.TxBytes), utils.DecodeBase64(data.Signature), config.MgoDevnet))

}
