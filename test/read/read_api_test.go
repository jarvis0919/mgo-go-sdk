package read

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

func TestGetChainIdentifier(t *testing.T) {
	identifier, err := devCli.MgoGetChainIdentifier(ctx)
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(identifier)
}

func TestGetCheckPoint(t *testing.T) {
	checkpoint, err := devCli.MgoGetCheckpoint(ctx, request.MgoGetCheckpointRequest{
		CheckpointID: "GJjR94qoW5fbDN4pJJMDBLiAZnyYKAKoDhzyYYBW8GPJ",
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(checkpoint)
}

func TestGetCheckPoints(t *testing.T) {
	checkpoints, err := devCli.MgoGetCheckpoints(ctx, request.MgoGetCheckpointsRequest{
		Cursor:          "18847201",
		Limit:           3,
		DescendingOrder: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(checkpoints)
}

func TestGetLatestCheckpointSequenceNumber(t *testing.T) {
	number, err := devCli.MgoGetLatestCheckpointSequenceNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(number)
}

func TestGetLoadedChildObjects(t *testing.T) {
	ob, err := devCli.MgoGetLoadedChildObjects(ctx, request.MgoGetLoadedChildObjectsRequest{
		Digest: "GxRMedP6tgLbwMY3fpUDgXMaRqZgw9ELhghFNVZSXLaU",
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(ob)
}

func TestGetObject(t *testing.T) {
	ob, err := devCli.MgoGetObject(ctx, request.MgoGetObjectRequest{
		ObjectId: "0x11ac113ffd2befec14988aa242635b3a59e2675bf11d95c07d055513bcbf6484",
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(ob)
}

func TestGetProtocolConfig(t *testing.T) {
	config, err := devCli.MgoGetProtocolConfig(ctx, request.MgoGetProtocolConfigRequest{
		Version: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(config)
}

func TestGetTotalTransactionBlocks(t *testing.T) {
	blocks, err := devCli.MgoGetTotalTransactionBlocks(ctx)
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(blocks)
}

func TestGetTransactionBlock(t *testing.T) {
	block, err := devCli.MgoGetTransactionBlock(ctx, request.MgoGetTransactionBlockRequest{
		Digest: "Ed7ZdSUJUbKZDnSQ1uGuEhZ85sR9Mh9xoAK89CyLh8CB",
	})
	if err != nil {
		t.Fatal(err)
	}
	utils.JsonPrint(block)
}
