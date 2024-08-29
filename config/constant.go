package config

import "math"

var (
	MGO_PRIVATE_KEY_PREFIX = "mgoprivkey"
	PRIVATE_KEY_SIZE       = 32
	MGO_ADDRESS_LENGTH     = 64
)

var (
	SIGNATURE_FLAG_TO_SCHEME = map[Scheme]string{
		0x00: "ED25519",
		0x01: "Secp256k1",
		// Add other schemes if applicable
	}
	SIGNATURE_SCHEME_TO_FLAG = map[string]Scheme{
		"ED25519":   0x00,
		"Secp256k1": 0x01,
		// Add other schemes if applicable
	}
	SIGNATURE_SCHEME_TO_SIZE = map[string]int{
		"ED25519":   32,
		"Secp256k1": 33,
	}
	DERIVATION_PATH_DEVNET = map[Scheme]string{
		0x00: `m/44'/784'/0'/0'/0'`,
		0x01: `m/54'/784'/0'/0/0`,
	}
	DERIVATION_PATH_TESTNET = map[Scheme]string{
		0x00: `m/44'/938'/0'/0'/0'`,
		0x01: `m/54'/938'/0'/0/0`,
	}
)

type Scheme byte
type Keytype byte

const (
	Ed25519Flag   Scheme = 0
	Secp256k1Flag Scheme = 1
	ErrorFlag     byte   = math.MaxUint8
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

	WssMgoDevnetEndpoint  = "wss://fullnode.devnet.mangonetwork.io"
	WssMgoTestnetEndpoint = "wss://fullnode.testnet.mangonetwork.io"
	WssMgoMainnetEndpoint = "wss://fullnode.mainnet.mangonetwork.io"

	FaucetTestnetEndpoint = ""
	FaucetDevnetEndpoint  = ""
)

type NetIdentity string

const (
	MgoDevnet  NetIdentity = "devnet"
	MgoTestnet NetIdentity = "testnet"
)
