package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jarvis0919/mgo-go-sdk/model"
	"github.com/jarvis0919/mgo-go-sdk/model/request"

	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"

	"github.com/tidwall/gjson"
)

func (c *Client) MgoCall(ctx context.Context, method string, params ...interface{}) (interface{}, error) {
	resp, err := c.conn.Request(ctx, httpconn.Operation{
		Method: method,
		Params: params,
	})
	if err != nil {
		return nil, err
	}
	if gjson.ParseBytes(resp).Get("error").Exists() {
		return nil, errors.New(gjson.ParseBytes(resp).Get("error").String())
	}
	return gjson.ParseBytes(resp).String(), nil
}

func (c *Client) MoveCall(ctx context.Context, req request.MoveCallRequest) (*model.TxnMetaData, error) {
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_moveCall",
		Params: []interface{}{
			req.Signer,
			req.PackageObjectId,
			req.Module,
			req.Function,
			req.TypeArguments,
			req.Arguments,
			req.Gas,
			req.GasBudget,
		},
	})
	if err != nil {
		return nil, err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return nil, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	var tx model.TxnMetaData
	err = json.Unmarshal([]byte(gjson.ParseBytes(respBytes).Get("result").String()), &tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
