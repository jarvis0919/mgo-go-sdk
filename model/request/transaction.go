package request

import (
	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/model"
)

type TransferMgoRequest struct {
	// the transaction signer's Mgo address
	Signer      signer.Signer `json:"signer"      yaml:"signer"`
	MgoObjectId string        `json:"mgoObjectId" yaml:"mgoObjectId"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget" yaml:"gasBudget"`
	Recipient string `json:"recipient" yaml:"recipient"`
	Amount    string `json:"amount"    yaml:"amount"`
}

type SignAndExecuteTransactionBlockRequest struct {
	TxnMetaData model.TxnMetaData
	// the address private key to sign the transaction
	Signer  signer.Signer
	Options TransactionBlockOptions `json:"options" yaml:"options"`
	// The optional enumeration values are: `WaitForEffectsCert`, or `WaitForLocalExecution`
	RequestType string `json:"requestType" yaml:"requestType"`
}
type TransactionBlockOptions struct {
	ShowInput          bool `json:"showInput,omitempty"          yaml:"showInput"`
	ShowRawInput       bool `json:"showRawInput,omitempty"       yaml:"showRawInput"`
	ShowEffects        bool `json:"showEffects,omitempty"        yaml:"showEffects"`
	ShowEvents         bool `json:"showEvents,omitempty"         yaml:"showEvents"`
	ShowObjectChanges  bool `json:"showObjectChanges,omitempty"  yaml:"showObjectChanges"`
	ShowBalanceChanges bool `json:"showBalanceChanges,omitempty" yaml:"showBalanceChanges"`
}
type MoveCallRequest struct {
	// the transaction signer's Mgo address
	Signer string `json:"signer"`
	// the package containing the module and function
	PackageObjectId string `json:"packageObjectId"`
	// the specific module in the package containing the function
	Module string `json:"module"`
	// the function to be called
	Function string `json:"function"`
	// the type arguments to the function
	TypeArguments []interface{} `json:"typeArguments"`
	// the arguments to the function
	Arguments []interface{} `json:"arguments"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas *string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`

	ExecutionMode string `json:"executionMode"`
}

type MgoSubscribeEventsRequest struct {
	// the event query criteria.
	MgoEventFilter interface{} `json:"mgoEventFilter"`
}
type MgoSubscribeTransactionsRequest struct {
	// the transaction query criteria.
	TransactionFilter interface{} `json:"filter"`
}
