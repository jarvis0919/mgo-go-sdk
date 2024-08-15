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

// MgoGetChainIdentifier implements the method `mgo_getChainIdentifier`, return the chain's identifier.
func (c *Client) MgoGetChainIdentifier(ctx context.Context) (string, error) {
	var rsp string
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgo_getChainIdentifier",
		Params: []interface{}{},
	})
	if err != nil {
		return rsp, err
	}

	if gjson.ParseBytes(respBytes).Get("error").Exists() {
		return rsp, errors.New(gjson.ParseBytes(respBytes).Get("error").String())
	}
	rsp = gjson.ParseBytes(respBytes).Get("result").String()
	return rsp, nil
}

// MgoGetCheckpoint implements the method `mgo_getCheckpoint`, gets a checkpoint.
func (c *Client) MgoGetCheckpoint(ctx context.Context, req request.MgoGetCheckpointRequest) (response.CheckpointResponse, error) {
	var rsp response.CheckpointResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgo_getCheckpoint",
		Params: []interface{}{
			req.CheckpointID,
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

// MgoGetCheckpoints implements the method `mgo_getCheckpoints`, gets paginated list of checkpoints.
func (c *Client) MgoGetCheckpoints(ctx context.Context, req request.MgoGetCheckpointsRequest) (response.PaginatedCheckpointsResponse, error) {
	var rsp response.PaginatedCheckpointsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgo_getCheckpoints",
		Params: []interface{}{
			req.Cursor,
			req.Limit,
			req.DescendingOrder,
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
