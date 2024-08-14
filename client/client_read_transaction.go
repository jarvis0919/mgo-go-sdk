package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/model/respone"
	"github.com/tidwall/gjson"
)

func (c *Client) MgoDevInspectTransactionBlock(ctx context.Context, req request.MgoDevInspectTransactionBlockRequest) (respone.MgoTransactionBlockResponse, error) {
	var rsp respone.MgoTransactionBlockResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgo_devInspectTransactionBlock",
		Params: []interface{}{
			req.Sender,
			req.TxBytes,
			req.GasPrice,
			req.Epoch,
		},
	})
	if err != nil {
		return rsp, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").Raw), &rsp)
	if err != nil {
		return rsp, err
	}
	return rsp, nil
}
