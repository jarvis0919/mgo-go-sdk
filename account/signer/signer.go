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

// func (s *Signer) ToMgoAddress() {
// 	flag, ok := global.SIGNATURE_SCHEME_TO_FLAG[s.Scheme]
// 	if !ok {
// 		return
// 	}
// 	inputBytes := append([]byte{flag}, []byte(s.PublicKey)...)
// 	s.MgoAddress = hex.EncodeToString(utils.Keccak256(inputBytes))[:global.MGO_ADDRESS_LENGTH*2]
// }

// 注释使用方法性能更高,但为了维护性，采用了append的方式处理
// tmp := make([]byte, len(publicKey)+1)
// tmp[0] = scheme
// for i := range publicKey {
// 	tmp[i+1] = publicKey[i]
// }
