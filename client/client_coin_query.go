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

// MgoXGetAllBalance implements the method `suix_getAllBalances`, gets all Coin balances owned by an address.
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
