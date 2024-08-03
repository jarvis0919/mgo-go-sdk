package client

import (
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/global"
)

func (c *Client) GetSignerAddress(s signer.Signer) string {
	switch c.net {
	case global.MgoDevnet:
		return s.ToMgoAddressDevNet()
	case global.MgoTestnet:
		return s.ToMgoAddressTestNet()
	default:
		return s.ToMgoAddressDevNet()
	}
}
