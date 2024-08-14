package extend

import (
	"context"
	"fmt"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"testing"
)

var ctx = context.Background()
var devCli = client.NewMgoClient(global.MgoDevnet)

func TestResolveNameServiceAddress(t *testing.T) {
	address, err := devCli.MgoXResolveNameServiceAddress(ctx, request.MgoXResolveNameServiceAddressRequest{
		Name: "example.mgo",
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	t.Log(address)
}

func TestResolveNameServiceNames(t *testing.T) {
	address, err := devCli.MgoXResolveNameServiceNames(ctx, request.MgoXResolveNameServiceNamesRequest{
		Address: "0x0000000000000000000000000000000000000000000000000000000000000002",
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	t.Log(address)
}
