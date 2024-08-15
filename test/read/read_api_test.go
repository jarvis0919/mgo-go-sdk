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
