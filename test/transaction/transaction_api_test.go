package transaction

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"os"
	"testing"
)

var ctx = context.Background()
var devCli = client.NewMgoClient(global.MgoDevnet)

func getSigner() (*signer.SignerEd25519, error) {
	// 文件中的私钥字符串为  ['private_key1','private_key2']
	bytes, err := os.ReadFile("../../private_keys.json")
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
func TestMergeCoin(t *testing.T) {
	mergeCoins, err := devCli.MergeCoins(ctx, request.MergeCoinsRequest{
		Signer:      "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		PrimaryCoin: "0x05678c9529d3354a291fc3235f445dc480ebd476fc281654e4731d2739a5e542",
		CoinToMerge: "0xb421a6f124cc4da9d12b4242a24eeb4be7d6e69871f53cf24ffe9deb35f66ccf",
		Gas:         "0x822f6705df64d073cbfeb2b2ef088f281aa2d486ea9b5c7fbb0ded58171d7f84",
		GasBudget:   "10000000",
	})
	if err != nil {
		t.Fatal(err)
	}
	ed25519Signer, err := getSigner()
	if err != nil {
		t.Fatal(err)
	}
	executeRes, err := devCli.SignAndExecuteTransactionBlock(ctx, request.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: mergeCoins,
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
	t.Log(executeRes)
}

func TestSplitCoin(t *testing.T) {
	splitCoins, err := devCli.SplitCoin(ctx, request.SplitCoinRequest{
		Signer:       "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		CoinObjectId: "0x91d2925ccb7be261e9db6f23daf9678a38945cb274014e1521b7441fbbc1a18d",
		SplitAmounts: []string{"1000", "1000"},
		Gas:          "0x9e9944e470b44c1363409505ef6d154562572a97cbca88dccfd0d972858b54a5",
		GasBudget:    "10000000",
	})
	if err != nil {
		t.Fatal(err)
	}
	ed25519Signer, err := getSigner()
	if err != nil {
		t.Fatal(err)
	}
	executeRes, err := devCli.SignAndExecuteTransactionBlock(ctx, request.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: splitCoins,
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
	t.Log(executeRes)
}

func TestSplitCoinEqual(t *testing.T) {
	splitCoins, err := devCli.SplitCoinEqual(ctx, request.SplitCoinEqualRequest{
		Signer:       "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		CoinObjectId: "0x91d2925ccb7be261e9db6f23daf9678a38945cb274014e1521b7441fbbc1a18d",
		SplitCount:   "3",
		Gas:          "0x9e9944e470b44c1363409505ef6d154562572a97cbca88dccfd0d972858b54a5",
		GasBudget:    "10000000",
	})
	if err != nil {
		t.Fatal(err)
	}
	ed25519Signer, err := getSigner()
	if err != nil {
		t.Fatal(err)
	}
	executeRes, err := devCli.SignAndExecuteTransactionBlock(ctx, request.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: splitCoins,
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
	t.Log(executeRes)
}
