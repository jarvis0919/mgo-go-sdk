package client

import "github.com/jarvis0919/mgo-go-sdk/client/wsconn"

// IMgoWebsocketAPI defines the subscription API related interface, and then implement it by the WebsocketClient.
type IMgoWebsocketAPI interface {
	ISubscribeAPI
}

// WebsocketClient implements MgoWebsocketAPI related interfaces.
type WebsocketClient struct {
	ISubscribeAPI
}

// NewMgoWebsocketClient instantiates the WebsocketClient to call the methods of each module.
func NewMgoWebsocketClient(rpcUrl string) IMgoWebsocketAPI {
	conn := wsconn.NewWsConn(rpcUrl)
	return &WebsocketClient{
		ISubscribeAPI: &mgoSubscribeImpl{
			conn: conn,
		},
	}
}
