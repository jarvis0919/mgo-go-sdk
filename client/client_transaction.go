package client

import (
	"context"
	"encoding/json"
	"errors"
	"mgo-go-sdk/account/keypair"
	"mgo-go-sdk/client/httpconn"
	"mgo-go-sdk/model"
	"mgo-go-sdk/model/request"
	"mgo-go-sdk/model/respone"

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

func (s *Client) SignAndExecuteTransactionBlock(ctx context.Context, req request.SignAndExecuteTransactionBlockRequest) (respone.MgoTransactionBlockResponse, error) {
	var rsp respone.MgoTransactionBlockResponse
	signedTxn := keypair.SignSerializedSigWith(&req.TxnMetaData, req.Signer.PrivateKeyBytes(), s.net)
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
