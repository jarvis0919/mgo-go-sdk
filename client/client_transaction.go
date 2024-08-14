package client

import (
	"context"
	"crypto/ed25519"
	"encoding/json"
	"errors"

	"github.com/jarvis0919/mgo-go-sdk/account/keypair"
	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"
	"github.com/jarvis0919/mgo-go-sdk/model"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/model/response"

	"github.com/tidwall/gjson"
)

func (c *Client) TransferMgo(ctx context.Context, req request.TransferMgoRequest) (model.TxnMetaData, error) {
	var rsp model.TxnMetaData
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_transferMgo",
		Params: []interface{}{
			c.GetSignerAddress(req.Signer),
			req.MgoObjectId,
			req.GasBudget,
			req.Recipient,
			req.Amount,
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

func (c *Client) TransferObject(ctx context.Context, req request.TransferObjectRequest) (model.TxnMetaData, error) {
	var rsp model.TxnMetaData
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "unsafe_transferObject",
		Params: []interface{}{
			c.GetSignerAddress(req.Signer),
			req.ObjectId,
			req.Gas,
			req.GasBudget,
			req.Recipient,
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

func (s *Client) SignAndExecuteTransactionBlock(ctx context.Context, req request.SignAndExecuteTransactionBlockRequest) (response.MgoTransactionBlockResponse, error) {
	var rsp response.MgoTransactionBlockResponse
	signedTxn := keypair.SignSerializedSigWith(&req.TxnMetaData, ed25519.NewKeyFromSeed(req.Signer.PrivateKeyBytes()), s.net)
	respBytes, err := s.conn.Request(ctx, httpconn.Operation{
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
