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

// MgoXGetAllBalance implements the method `mgox_getAllBalances`, gets all Coin balances owned by an address.
func (c *Client) MgoXGetAllBalance(ctx context.Context, req request.MgoXGetAllBalanceRequest) (response.CoinAllBalanceResponse, error) {
	var rsp response.CoinAllBalanceResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_getAllBalances",
		Params: []interface{}{
			req.Owner,
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

// MgoXGetAllCoins implements the method `mgox_getAllCoins`, gets all Coin objects owned by an address.
func (c *Client) MgoXGetAllCoins(ctx context.Context, req request.MgoXGetAllCoinsRequest) (response.PaginatedCoinsResponse, error) {
	var rsp response.PaginatedCoinsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_getAllCoins",
		Params: []interface{}{
			req.Owner,
			req.Cursor,
			req.Limit,
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

// MgoXGetCoinMetadata implements the method `mgox_getCoinMetadata`, gets metadata(e.g., symbol, decimals) for a coin.
func (c *Client) MgoXGetCoinMetadata(ctx context.Context, req request.MgoXGetCoinMetadataRequest) (response.CoinMetadataResponse, error) {
	var rsp response.CoinMetadataResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_getCoinMetadata",
		Params: []interface{}{
			req.CoinType,
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
