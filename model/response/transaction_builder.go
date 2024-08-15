package response

import "github.com/jarvis0919/mgo-go-sdk/model/mgo_types"

type BatchTransactionResponse struct {
	Gas          []mgo_types.MgoObjectRef `json:"gas"`
	InputObjects []interface{}            `json:"inputObjects"`
	TxBytes      string                   `json:"txBytes"`
}
