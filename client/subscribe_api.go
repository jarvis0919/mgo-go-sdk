package client

import (
	"context"
	"encoding/json"
	"github.com/jarvis0919/mgo-go-sdk/client/wsconn"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/jarvis0919/mgo-go-sdk/model/response"
	"github.com/tidwall/gjson"
	"log"
)

type ISubscribeAPI interface {
	SubscribeEvent(ctx context.Context, req request.MgoSubscribeEventsRequest, msgCh chan response.MgoEventResponse) error
	SubscribeTransaction(ctx context.Context, req request.MgoSubscribeTransactionsRequest, msgCh chan response.MgoEffects) error
}
type mgoSubscribeImpl struct {
	conn *wsconn.WsConn
}

// SubscribeEvent implements the method `mgox_subscribeEvent`, subscribe to a stream of Mgo event.
func (s *mgoSubscribeImpl) SubscribeEvent(ctx context.Context, req request.MgoSubscribeEventsRequest, msgCh chan response.MgoEventResponse) error {
	rsp := make(chan []byte, 10)
	err := s.conn.Call(ctx, wsconn.CallOp{
		Method: "mgox_subscribeEvent",
		Params: []interface{}{
			req.MgoEventFilter,
		},
	}, rsp)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case messageData := <-rsp:
				var result response.MgoEventResponse
				if gjson.ParseBytes(messageData).Get("error").Exists() {
					log.Fatal(gjson.ParseBytes(messageData).Get("error").String())
				}

				err := json.Unmarshal([]byte(gjson.ParseBytes(messageData).Get("params.result").String()), &result)
				if err != nil {
					log.Fatal(err)
				}

				msgCh <- result
			}
		}
	}()

	return nil
}

// SubscribeTransaction implements the method `mgox_subscribeTransaction`, subscribe to a stream of Mgo transaction effects.
func (s *mgoSubscribeImpl) SubscribeTransaction(ctx context.Context, req request.MgoSubscribeTransactionsRequest, msgCh chan response.MgoEffects) error {
	rsp := make(chan []byte, 10)
	err := s.conn.Call(ctx, wsconn.CallOp{
		Method: "mgox_subscribeTransaction",
		Params: []interface{}{
			req.TransactionFilter,
		},
	}, rsp)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case messageData := <-rsp:
				var result response.MgoEffects
				if gjson.ParseBytes(messageData).Get("error").Exists() {
					log.Fatal(gjson.ParseBytes(messageData).Get("error").String())
				}

				err := json.Unmarshal([]byte(gjson.ParseBytes(messageData).Get("params.result").String()), &result)
				if err != nil {
					log.Fatal(err)
				}

				msgCh <- result
			}
		}
	}()

	return nil
}
