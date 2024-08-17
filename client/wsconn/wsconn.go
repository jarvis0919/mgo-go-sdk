package wsconn

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jarvis0919/mgo-go-sdk/model/request"
	"github.com/tidwall/gjson"
	"log"
	"time"
)

type WsConn struct {
	Conn   *websocket.Conn
	wsUrl  string
	ticker *time.Ticker // For heartbeat, default 30s
}

type CallOp struct {
	Method string
	Params []interface{}
}

func NewWsConn(wsUrl string) *WsConn {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsUrl, nil)

	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err, wsUrl)
	}

	return &WsConn{
		Conn:   conn,
		wsUrl:  wsUrl,
		ticker: time.NewTicker(30 * time.Second),
	}
}

func NewWsConnWithDuration(wsUrl string, d time.Duration) *WsConn {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsUrl, nil)

	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err, wsUrl)
	}

	return &WsConn{
		Conn:   conn,
		wsUrl:  wsUrl,
		ticker: time.NewTicker(d),
	}
}

func (w *WsConn) Call(ctx context.Context, op CallOp, receiveMsgCh chan []byte) error {
	jsonRPCCall := request.JsonRPCRequest{
		JsonRPC: "2.0",
		ID:      time.Now().UnixMilli(),
		Method:  op.Method,
		Params:  op.Params,
	}

	callBytes, err := json.Marshal(jsonRPCCall)
	if err != nil {
		return err
	}

	err = w.Conn.WriteMessage(websocket.TextMessage, callBytes)
	if nil != err {
		return err
	}

	_, messageData, err := w.Conn.ReadMessage()
	if nil != err {
		return err
	}

	var rsp SubscriptionResp
	if gjson.ParseBytes(messageData).Get("error").Exists() {
		return fmt.Errorf(gjson.ParseBytes(messageData).Get("error").String())
	}

	err = json.Unmarshal([]byte(gjson.ParseBytes(messageData).String()), &rsp)
	if err != nil {
		return err
	}

	fmt.Printf("establish successfully, subscriptionID: %d, Waiting to accept data...\n", rsp.Result)

	// Start a ticker for sending pings
	go func() {
		for {
			select {
			case <-ctx.Done():
				w.ticker.Stop()
				return
			case <-w.ticker.C:
				// Send a ping
				if err := w.Conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
					log.Println("ping failed:", err)
					return
				} else {
					log.Println("heartbeat...")
				}
			}
		}
	}()

	go func(conn *websocket.Conn) {
		for {
			messageType, messageData, err := conn.ReadMessage()
			if nil != err {
				log.Println(err)
				break
			}
			switch messageType {
			case websocket.TextMessage:
				receiveMsgCh <- messageData

			default:
				continue
			}
		}
	}(w.Conn)

	return nil
}
