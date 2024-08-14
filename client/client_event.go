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

// MgoGetEvents implements the method `mgo_getEvents`, gets transaction events.
func (c *Client) MgoGetEvents(ctx context.Context, req request.MgoGetEventsRequest) (response.GetEventsResponse, error) {
	var rsp response.GetEventsResponse
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgo_getEvents",
		Params: []interface{}{
			req.Digest,
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

// MgoXQueryEvents implements the method `mgox_queryEvents`, gets list of events for a specified query criteria.
func (c *Client) MgoXQueryEvents(ctx context.Context, req request.MgoXQueryEventsRequest) (response.PaginatedEventsResponse, error) {
	var rsp response.PaginatedEventsResponse
	if err := validate.ValidateStruct(req); err != nil {
		return rsp, err
	}
	respBytes, err := c.conn.Request(ctx, httpconn.Operation{
		Method: "mgox_queryEvents",
		Params: []interface{}{
			req.MgoEventFilter,
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
