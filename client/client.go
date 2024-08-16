package client

import (
	"github.com/jarvis0919/mgo-go-sdk/client/httpconn"
	"github.com/jarvis0919/mgo-go-sdk/config"
)

type Client struct {
	conn *httpconn.HttpConn
	net  config.NetIdentity
}

func NewMgoClient(net config.NetIdentity) *Client {
	rpcUrl := config.RPC_MGO_NET_URL[net]
	conn := httpconn.NewHttpConn(rpcUrl)
	return newClient(conn, net)
}
func newClient(conn *httpconn.HttpConn, net config.NetIdentity) *Client {
	return &Client{
		conn: conn,
		net:  net,
	}
}
