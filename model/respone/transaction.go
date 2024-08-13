package respone

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
type EventId struct {
	TxDigest string `json:"txDigest"`
	EventSeq string `json:"eventSeq"`
}
type MgoEventResponse struct {
	Id                EventId                `json:"id"`
	PackageId         string                 `json:"packageId"`
	TransactionModule string                 `json:"transactionModule"`
	Sender            string                 `json:"sender"`
	Type              string                 `json:"type"`
	ParsedJson        map[string]interface{} `json:"parsedJson"`
	Bcs               string                 `json:"bcs"`
	TimestampMs       string                 `json:"timestampMs"`
}

type MgoEffects struct {
	MessageVersion     string               `json:"messageVersion"`
	Status             ExecutionStatus      `json:"status"`
	ExecutedEpoch      string               `json:"executedEpoch"`
	GasUsed            GasCostSummary       `json:"gasUsed"`
	ModifiedAtVersions []ModifiedAtVersions `json:"modifiedAtVersions"`
	SharedObjects      []MgoObjectRef       `json:"sharedObjects"`
	TransactionDigest  string               `json:"transactionDigest"`
	Created            []OwnedObjectRef     `json:"created"`
	Mutated            []OwnedObjectRef     `json:"mutated"`
	Deleted            []MgoObjectRef       `json:"deleted"`
	GasObject          OwnedObjectRef       `json:"gasObject"`
	EventsDigest       string               `json:"eventsDigest"`
	Dependencies       []string             `json:"dependencies"`
}
type ExecutionStatus struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type GasCostSummary struct {
	ComputationCost         string `json:"computationCost"`
	StorageCost             string `json:"storageCost"`
	StorageRebate           string `json:"storageRebate"`
	NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
}
type ModifiedAtVersions struct {
	ObjectId       string `json:"objectId"`
	SequenceNumber string `json:"sequenceNumber"`
}
type MgoObjectRef struct {
	ObjectId string `json:"objectId"`
	Version  int    `json:"version"`
	Digest   string `json:"digest"`
}
type OwnedObjectRef struct {
	Owner     interface{}  `json:"owner"`
	Reference MgoObjectRef `json:"reference"`
}
