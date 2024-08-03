package respone

import "mgo-go-sdk/model"

type MgoTransactionBlockResponse struct {
	Digest                  string                 `json:"digest"`
	Transaction             model.TransactionBlock `json:"transaction,omitempty"`
	RawTransaction          string                 `json:"rawTransaction,omitempty"`
	Effects                 model.Effects          `json:"effects,omitempty"`
	Events                  []model.EventResponse  `json:"events,omitempty"`
	ObjectChanges           []model.ObjectChange   `json:"objectChanges,omitempty"`
	BalanceChanges          []model.BalanceChanges `json:"balanceChanges,omitempty"`
	TimestampMs             string                 `json:"timestampMs,omitempty"`
	Checkpoint              string                 `json:"checkpoint,omitempty"`
	ConfirmedLocalExecution bool                   `json:"confirmedLocalExecution,omitempty"`
}
