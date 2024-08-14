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

// MgoXResolveNameServiceAddress implements the method `mgox_resolveNameServiceAddress`, get the resolved address given resolver and name.
func (c *Client) MgoXResolveNameServiceAddress(ctx context.Context, req request.MgoXResolveNameServiceAddressRequest) (string, error) {
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_resolveNameServiceAddress",
		Params: []interface{}{
			req.Name,
		},
	})
	if err != nil {
		return "", err
	}
	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return "", errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}

	return gjson.ParseBytes(respBytes).Get("result").String(), nil
}

// MgoXResolveNameServiceNames implements the method `mgox_resolveNameServiceNames`, return the resolved names given address, if multiple names are resolved, the first one is the primary name.
func (c *Client) MgoXResolveNameServiceNames(ctx context.Context, req request.MgoXResolveNameServiceNamesRequest) (response.MgoXResolveNameServiceNamesResponse, error) {
	var rsp response.MgoXResolveNameServiceNamesResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_resolveNameServiceNames",
		Params: []interface{}{
			req.Address,
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
