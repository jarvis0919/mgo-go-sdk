package request

type RPCTransactionRequestParams struct {
	MoveCallRequestParams       *MoveCallRequest       `json:"moveCallRequestParams,omitempty"`
	TransferObjectRequestParams *TransferObjectRequest `json:"transferObjectRequestParams,omitempty"`
}

type BatchTransactionRequest struct {
	// the transaction signer's Mgo address
	Signer string `json:"signer"`
	// list of transaction request parameters
	RPCTransactionRequestParams []RPCTransactionRequestParams `json:"RPCTransactionRequestParams"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
	// Whether this is a regular transaction or a Dev Inspect Transaction
	// The optional enumeration values are: `DevInspect`, or `Commit`
	MgoTransactionBlockBuilderMode string `json:"mgoTransactionBlockBuilderMode"`
}

type PayRequest struct {
	// the transaction signer's Mgo address
	Signer      string   `json:"signer"`
	MgoObjectId []string `json:"mgoObjectId"`
	Recipient   []string `json:"recipient"`
	Amount      []string `json:"amount"`
	// gas object to be used in this transaction, node will pick one from the signer's possession if not provided
	Gas string `json:"gas"`
	// the gas budget, the transaction will fail if the gas cost exceed the budget
	GasBudget string `json:"gasBudget"`
}
