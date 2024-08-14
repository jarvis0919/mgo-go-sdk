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

// MgoXGetReferenceGasPrice implements the method `mgox_getReferenceGasPrice`, gets the reference gas price for the network.
func (c *Client) MgoXGetReferenceGasPrice(ctx context.Context) (uint64, error) {
	var rsp uint64
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_getReferenceGasPrice",
		Params: []interface{}{},
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

// MgoXGetStakes implements the method `mgox_getStakes`, gets the delegated stakes for an address.
func (c *Client) MgoXGetStakes(ctx context.Context, req request.MgoXGetStakesRequest) ([]*response.DelegatedStakesResponse, error) {
	var rsp []*response.DelegatedStakesResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_getStakes",
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

// MgoXGetStakesByIds implements the method `mgox_getStakesByIds`, return one or more delegated stake. If a Stake was withdrawn, its status will be Unstaked.
func (c *Client) MgoXGetStakesByIds(ctx context.Context, req request.MgoXGetStakesByIdsRequest) ([]*response.DelegatedStakesResponse, error) {
	var rsp []*response.DelegatedStakesResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_getStakesByIds",
		Params: []interface{}{
			req.StakedMgoIds,
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
