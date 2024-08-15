package extend

import (
	"context"
	"fmt"
	"github.com/jarvis0919/mgo-go-sdk/client"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/utils"
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

func TestGetDynamicFieldObject(t *testing.T) {
	object, err := devCli.MgoXGetDynamicFieldObject(ctx, request.MgoXGetDynamicFieldObjectRequest{
		ObjectId: "0x11ac113ffd2befec14988aa242635b3a59e2675bf11d95c07d055513bcbf6484",
		DynamicFieldName: request.DynamicFieldObjectName{
			Type:  "0x2::mgp::MGO",
			Value: "",
		},
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	t.Log(object)
}

func TestGetDynamicFields(t *testing.T) {
	object, err := devCli.MgoXGetDynamicFields(ctx, request.MgoXGetDynamicFieldsRequest{
		ObjectId: "0x11ac113ffd2befec14988aa242635b3a59e2675bf11d95c07d055513bcbf6484",
		Cursor:   "0xa9334aeacc435c70ab9635e47a277d8f8dd9d87765d1aadec2db8cc24c312542",
		Limit:    3,
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	t.Log(object)
}

func TestOwnedObjects(t *testing.T) {
	object, err := devCli.MgoXGetOwnedObjects(ctx, request.MgoXGetOwnedObjectsRequest{
		Address: "0x6d5ae691047b8e55cb3fc84da59651c5bae57d2970087038c196ed501e00697b",
		Limit:   10,
	})
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	utils.JsonPrint(object)
}
