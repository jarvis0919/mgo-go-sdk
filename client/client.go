package client

import (
	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"
	"github.com/jarvis0919/mgo-go-sdk/global"
)

type Client struct {
	conn *httpconn.HttpConn
	net  global.NetIdentity
}

func NewMgoClient(net global.NetIdentity) *Client {
	rpcUrl := global.RPC_MGO_NET_URL[net]
	conn := httpconn.NewHttpConn(rpcUrl)
	return newClient(conn, net)
}
func newClient(conn *httpconn.HttpConn, net global.NetIdentity) *Client {
	return &Client{
		conn: conn,
		net:  net,
	}
}
