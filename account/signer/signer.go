package signer

type Signer interface {
	Sign(message []byte) []byte

	SecretKeyHex() string
	PublicKeyHex() string
	PrivateKeyHex() string

	SecretKeyBytes() []byte
	PublicKeyBytes() []byte
	PrivateKeyBytes() []byte

	// MgoPrivateKey() string

	ToMgoAddressDevNet() string
	ToMgoAddressTestNet() string
}
