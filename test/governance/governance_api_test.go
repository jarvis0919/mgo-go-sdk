package governance

import (
	"context"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"testing"
)

var ctx = context.Background()
var devCli = client.NewMgoClient(global.MgoDevnet)

func TestGetReferenceGasPrice(t *testing.T) {
	price, err := devCli.MgoXGetReferenceGasPrice(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(price)
}
