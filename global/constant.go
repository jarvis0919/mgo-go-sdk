package global

import "math"

var (
	MGO_PRIVATE_KEY_PREFIX = "mgoprivkey"
	PRIVATE_KEY_SIZE       = 32
	MGO_ADDRESS_LENGTH     = 64
)

var (
	SIGNATURE_FLAG_TO_SCHEME = map[Keytype]string{
		0x00: "Ed25519",
		0x01: "Secp256k1",
		// Add other schemes if applicable
	}
	SIGNATURE_SCHEME_TO_FLAG = map[string]Keytype{
		"Ed25519":   0x00,
		"Secp256k1": 0x01,
		// Add other schemes if applicable
	}
	DERIVATION_PATH_DEVNET = map[Keytype]string{
		0x00: `m/44'/784'/0'/0'/0'`,
		0x01: `m/54'/784'/0'/0/0`,
	}
	DERIVATION_PATH_TESTNET = map[Keytype]string{
		0x00: `m/44'/938'/0'/0'/0'`,
		0x01: `m/54'/938'/0'/0/0`,
	}
)

type Keytype byte

const (
	Ed25519Flag   Keytype = 0
	Secp256k1Flag Keytype = 1
	ErrorFlag     byte    = math.MaxUint8
)
const (
	TransactionData    Keytype = 0
	TransactionEffects Keytype = 1
	CheckpointSummary  Keytype = 2
	PersonalMessage    Keytype = 3
)
const (
	Ed25519PublicKeyLength   = 32
	Secp256k1PublicKeyLength = 33
)

const (
	DefaultAccountAddressLength = 16
	AccountAddress20Length      = 20
	AccountAddress32Length      = 32
)

var RPC_MGO_NET_URL = map[NetIdentity]string{
	"testnet": RpcMgoTestnetEndpoint,
	"devnet":  RpcMgoDevnetEndpoint,
}

const (
	RpcMgoTestnetEndpoint = "https://indexer.testnet.mangonetwork.io/"
	RpcMgoDevnetEndpoint  = "https://indexer.devnet.mangonetwork.io/"

	WssMgoDevnetEndpoint  = ""
	WssMgoTestnetEndpoint = ""
	WssMgoMainnetEndpoint = ""

	FaucetTestnetEndpoint = ""
	FaucetDevnetEndpoint  = ""
)

type NetIdentity string

const (
	MgoDevnet  NetIdentity = "devnet"
	MgoTestnet NetIdentity = "testnet"
)
