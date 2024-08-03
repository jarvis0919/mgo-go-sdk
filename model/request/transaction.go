package request

import (
	"mgo-go-sdk/account/signer"
	"mgo-go-sdk/model"
)

type TransferMgoRequest struct {
	// the transaction signer's Mgo address
	Signer      signer.Signer `json:"signer"`
	MgoObjectId string        `json:"mgoObjectId"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

type SignAndExecuteTransactionBlockRequest struct {
	TxnMetaData model.TxnMetaData
	// the address private key to sign the transaction
	Signer  signer.Signer
	Options TransactionBlockOptions `json:"options"`
	// The optional enumeration values are: `WaitForEffectsCert`, or `WaitForLocalExecution`
	RequestType string `json:"requestType"`
}
type TransactionBlockOptions struct {
	ShowInput          bool `json:"showInput,omitempty"`
	ShowRawInput       bool `json:"showRawInput,omitempty"`
	ShowEffects        bool `json:"showEffects,omitempty"`
	ShowEvents         bool `json:"showEvents,omitempty"`
	ShowObjectChanges  bool `json:"showObjectChanges,omitempty"`
	ShowBalanceChanges bool `json:"showBalanceChanges,omitempty"`
}
