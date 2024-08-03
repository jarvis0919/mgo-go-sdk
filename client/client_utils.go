package client

import (
	"mgo-go-sdk/account/signer"
	"mgo-go-sdk/global"
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
