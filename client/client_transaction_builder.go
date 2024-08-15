package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"
	"github.com/jarvis0919/mgo-go-sdk/model"
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

// Pay implements the method `unsafe_pay`, send `Coin<T>` to a list of addresses, where `T` can be any coin type, following a list of amounts.
// The object specified in the `gas` field will be used to pay the gas fee for the transaction.
// The gas object can not appear in `input_coins`. If the gas object is not specified, the RPC server will auto-select one.
func (c *Client) Pay(ctx context.Context, req request.PayRequest) (model.TxnMetaData, error) {
	var rsp model.TxnMetaData
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_pay",
		Params: []interface{}{
			req.Signer,
			req.MgoObjectId,
			req.Recipient,
			req.Amount,
			req.Gas,
			req.GasBudget,
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

// PayAllMgo implements the method `unsafe_payAllMgo`, send all SUI coins to one recipient.
// This is for SUI coin only and does not require a separate gas coin object.
// Specifically, what pay_all_mgo does are:
// 1. accumulate all SUI from input coins and deposit all SUI to the first input coin.
// 2. transfer the updated first coin to the recipient and also use this first coin as gas coin object.
// 3. the balance of the first input coin after tx is sum(input_coins) - actual_gas_cost.
// 4. all other input coins other than the first are deleted.
func (c *Client) PayAllMgo(ctx context.Context, req request.PayAllMgoRequest) (model.TxnMetaData, error) {
	var rsp model.TxnMetaData
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_payAllMgo",
		Params: []interface{}{
			req.Signer,
			req.MgoObjectId,
			req.Recipient,
			req.GasBudget,
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
