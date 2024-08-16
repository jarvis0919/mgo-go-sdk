package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/model/response"

	"github.com/tidwall/gjson"
)

func (c *Client) SignAndExecuteTransactionBlock(ctx context.Context, req request.SignAndExecuteTransactionBlockRequest) (response.MgoTransactionBlockResponse, error) {
	var rsp response.MgoTransactionBlockResponse
	signedTxn := req.Keypair.SignTransactionBlock(&req.TxnMetaData, c.net)
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgo_executeTransactionBlock",
		Params: []interface{}{
			signedTxn.TxBytes,
			[]string{signedTxn.Signature},
			req.Options,
			req.RequestType,
		},
	})

	if err != nil {
		return rsp, err
	}

	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}

	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &rsp)
	if err != nil {
		return rsp, err
	}

	return rsp, nil
}
