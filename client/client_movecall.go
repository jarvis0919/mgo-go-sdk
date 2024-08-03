package client

import (
	"context"
	"errors"

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
