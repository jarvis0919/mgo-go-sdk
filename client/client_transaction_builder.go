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

// BatchTransaction implements the method `unsafe_batchTransaction`, creates an unsigned batched transaction.
func (c *Client) BatchTransaction(ctx context.Context, req request.BatchTransactionRequest) (response.BatchTransactionResponse, error) {
	var rsp response.BatchTransactionResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_batchTransaction",
		Params: []interface{}{
			req.Signer,
			req.RPCTransactionRequestParams,
			req.Gas,
			req.GasBudget,
			req.MgoTransactionBlockBuilderMode,
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
