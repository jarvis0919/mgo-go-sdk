package transaction_builder

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/utils"
	"os"
	"testing"
)

var ctx = context.Background()
var devCli = client.NewMgoClient(global.MgoDevnet)

func TestBatchTransaction(t *testing.T) {
	rsp, err := devCli.BatchTransaction(ctx, request.BatchTransactionRequest{
		Signer: "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		RPCTransactionRequestParams: []request.RPCTransactionRequestParams{
			{
				MoveCallRequestParams: &request.MoveCallRequest{
					PackageObjectId: "0x0000000000000000000000000000000000000000000000000000000000000002",
					Module:          "mgo",
					Function:        "transfer",
					TypeArguments:   []interface{}{},
					Arguments: []interface{}{
						"0xe586e913e413de2df45e3eca0f0adf342a1f5d8d71e61805e14fd4872529f727",
						"0x0cafa361487490f306c0b4c3e4cf0dc6fd584c5259ab1d5457d80a9e2170e238",
					},
				},
			},
			{
				TransferObjectRequestParams: &request.TransferObjectRequest{
					ObjectId:  "0x11ac113ffd2befec14988aa242635b3a59e2675bf11d95c07d055513bcbf6484",
					Recipient: "0x0cafa361487490f306c0b4c3e4cf0dc6fd584c5259ab1d5457d80a9e2170e238",
				},
			},
		},
		Gas:                            "0xc7f7956990f4d210024f34d9c5692e0b42d01aa67c94532521e8be886a0dbaa7",
		GasBudget:                      "100000",
		MgoTransactionBlockBuilderMode: "DevInspect",
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(rsp)
}

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

func TestPay(t *testing.T) {
	pay, err := devCli.Pay(ctx, request.PayRequest{
		Signer:      "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		MgoObjectId: []string{"0x05678c9529d3354a291fc3235f445dc480ebd476fc281654e4731d2739a5e542", "0x11ac113ffd2befec14988aa242635b3a59e2675bf11d95c07d055513bcbf6484"},
		Recipient:   []string{"0xbb3888e6c078a8ccedde58394873584ba39878984f1f8da4cba870de7eb5c3d2", "0x0cafa361487490f306c0b4c3e4cf0dc6fd584c5259ab1d5457d80a9e2170e238"},
		Amount:      []string{"1000", "1000"},
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
		TxnMetaData: pay,
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

func TestPayAllMgo(t *testing.T) {
	pay, err := devCli.PayAllMgo(ctx, request.PayAllMgoRequest{
		Signer:      "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		MgoObjectId: []string{"0xe586e913e413de2df45e3eca0f0adf342a1f5d8d71e61805e14fd4872529f727", "0xc7f7956990f4d210024f34d9c5692e0b42d01aa67c94532521e8be886a0dbaa7"},
		Recipient:   "0x0cafa361487490f306c0b4c3e4cf0dc6fd584c5259ab1d5457d80a9e2170e238",
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
		TxnMetaData: pay,
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
