package client

import (
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/config"
)

func (c *Client) GetSignerAddress(s signer.Signer) string {
	switch c.net {
	case config.MgoDevnet:
		return s.ToMgoAddressDevNet()
	case config.MgoTestnet:
		return s.ToMgoAddressTestNet()
	default:
		return s.ToMgoAddressDevNet()
	}
}
