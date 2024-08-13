package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/utils"
	"os"
)

var ctx = context.Background()
var devCli = client.NewMgoClient(global.MgoDevnet)

func getMoveCallData() (*model.TxnMetaData, error) {
	gas := "0x9e9944e470b44c1363409505ef6d154562572a97cbca88dccfd0d972858b54a5"
	req := request.MoveCallRequest{
		Signer:          "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		PackageObjectId: "0x0000000000000000000000000000000000000000000000000000000000000002",
		Module:          "mgo",
		Function:        "transfer",
		TypeArguments:   []interface{}{},
		Arguments: []interface{}{
			"0x171e4c8a943fd30567a90a4d3293c06a3ebd317c5f4b8a119e942264ffa4e122",
			"0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		},
		Gas:       &gas,
		GasBudget: "1000",
	}
	return devCli.MoveCall(ctx, req)
}
func getSigner() (*signer.SignerEd25519, error) {
	// 文件中的私钥字符串为  ['private_key1','private_key2']
	bytes, err := os.ReadFile("private_keys.json")
	if err != nil {
		return nil, err
	}
	store := []string{}
	err = json.Unmarshal(bytes, &store)
	if err != nil {
		return nil, err
	}

	key, err := signer.NewEd25519SignerFromPrivateKey("0xa9c6efc5ffc3372f29b108b5ac039f3cf8d411b953b9d212f48b22c3620a5a56")
	if err != nil {
		return nil, err
	}
	return key, nil
}

func moveCall() {
	// 1. 先调用 moveCall 生成交易数据
	// 2. 使用私钥签署

	// 1.
	txnMetaData, err := getMoveCallData()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	utils.JsonPrint(txnMetaData)
	// 2.
	ed25519Signer, err := getSigner()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	executeRes, err := devCli.SignAndExecuteTransactionBlock(ctx, request.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: *txnMetaData,
		Signer:      ed25519Signer,
		// only fetch the effects field
		Options: request.TransactionBlockOptions{
			ShowInput:    true,
			ShowRawInput: true,
			ShowEffects:  true,
		},
		RequestType: "WaitForLocalExecution",
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(executeRes)

}
func main() {
	moveCall()
}
