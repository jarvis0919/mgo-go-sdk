package transaction_builder

import (
	"context"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/utils"
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
