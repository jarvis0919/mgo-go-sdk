package response

import "github.com/jarvis0919/mgo-go-sdk/model"

type MgoTransactionBlockResponse struct {
	Digest                  string                 `json:"digest"                            yaml:"digest"`
	Transaction             model.TransactionBlock `json:"transaction,omitempty"             yaml:"transaction"`
	RawTransaction          string                 `json:"rawTransaction,omitempty"          yaml:"rawTransaction"`
	Effects                 model.Effects          `json:"effects,omitempty"                 yaml:"effects"`
	Events                  []model.EventResponse  `json:"events,omitempty"                  yaml:"events"`
	ObjectChanges           []model.ObjectChange   `json:"objectChanges,omitempty"           yaml:"objectChanges"`
	BalanceChanges          []model.BalanceChanges `json:"balanceChanges,omitempty"          yaml:"balanceChanges"`
	TimestampMs             string                 `json:"timestampMs,omitempty"             yaml:"timestampMs"`
	Checkpoint              string                 `json:"checkpoint,omitempty"              yaml:"checkpoint"`
	ConfirmedLocalExecution bool                   `json:"confirmedLocalExecution,omitempty" yaml:"confirmedLocalExecution"`
}

type TransactionFilter map[string]interface{}

type MgoTransactionBlockOptions struct {
	ShowInput          bool `json:"showInput,omitempty"`
	ShowRawInput       bool `json:"showRawInput,omitempty"`
	ShowEffects        bool `json:"showEffects,omitempty"`
	ShowEvents         bool `json:"showEvents,omitempty"`
	ShowObjectChanges  bool `json:"showObjectChanges,omitempty"`
	ShowBalanceChanges bool `json:"showBalanceChanges,omitempty"`
}

type MgoTransactionBlockResponseQuery struct {
	TransactionFilter TransactionFilter          `json:"filter"`
	Options           MgoTransactionBlockOptions `json:"options"`
}
