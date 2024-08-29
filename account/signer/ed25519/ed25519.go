package ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/config"

	"github.com/jarvis0919/mgo-go-sdk/utils"
)

type SignerEd25519 struct {
	PrivateKey        ed25519.PrivateKey
	PublicKey         ed25519.PublicKey
	MgoAddressDevNet  string
	MgoAddressTestNet string
}

func NewEd25519Signer() (*SignerEd25519, error) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return newEd25519Signer(privateKey)
}
func NewEd25519SignerFromPrivateKey(privatekey string) (ed25519Signer *SignerEd25519, err error) {
	if strings.HasPrefix(privatekey, "0x") || strings.HasPrefix(privatekey, "0X") {
		privatekey = privatekey[2:]
	}
	seed, err := hex.DecodeString(privatekey)
	if err != nil {
		return nil, err
	}
	return newEd25519SignerFromSeed(seed)
}

func NewEd25519SignerFromMgoPrivatekey(mgoPrivatekey string) (ed25519Signer *SignerEd25519, err error) {
	ParsedKeypair, err := signer.DecodeMgoPrivateKey(mgoPrivatekey)
	if err != nil {
		return nil, err
	}
	return newEd25519SignerFromSeed(ParsedKeypair.SecretKey)
}

func newEd25519SignerFromSeed(seed []byte) (ed25519Signer *SignerEd25519, err error) {
	defer func() {
		if r := recover(); r != nil {
			ed25519Signer = nil
			err = fmt.Errorf("recovered from panic: %s", r)
		}
	}()
	return newEd25519Signer(ed25519.NewKeyFromSeed(seed))
}

func newEd25519Signer(privateKey ed25519.PrivateKey) (*SignerEd25519, error) {
	publicKey := privateKey.Public().(ed25519.PublicKey)
	return &SignerEd25519{
		PrivateKey:        privateKey,
		PublicKey:         publicKey,
		MgoAddressDevNet:  mgoAddressDevNet(publicKey),
		MgoAddressTestNet: mgoAddressTestNet(publicKey),
	}, nil
}

func (s SignerEd25519) String() string {
	return fmt.Sprintf("PrivateKey: %s\nPublicKey: %s\nMgoAddressDevNet: %s\nMgoAddressTestNet: %s",
		hex.EncodeToString(s.PrivateKey),
		hex.EncodeToString(s.PublicKey),
		s.MgoAddressDevNet, s.MgoAddressTestNet)
}

func (s *SignerEd25519) MgoPrivateKey() string {
	mgoPrivateKey, _ := signer.EncodeMgoPrivateKey(s.PrivateKey[:config.PRIVATE_KEY_SIZE], config.Ed25519Flag)
	return mgoPrivateKey
}

func (s *SignerEd25519) SecretKeyHex() string {
	return "0x" + hex.EncodeToString(s.PrivateKey)
}

func (s *SignerEd25519) SecretKeyBytes() []byte {
	return s.PrivateKey
}

func (s *SignerEd25519) PrivateKeyHex() string {
	return "0x" + hex.EncodeToString(s.PrivateKey[:32])
}

func (s *SignerEd25519) PrivateKeyBytes() []byte {
	return s.PrivateKey[:32]
}

func (s *SignerEd25519) PublicKeyHex() string {
	return "0x" + hex.EncodeToString(s.PublicKey)
}

func (s *SignerEd25519) PublicKeyBytes() []byte {
	return s.PublicKey
}

func (s *SignerEd25519) PublicBase64Key() string {
	return utils.EncodeBase64(s.PublicKey)
}
func (s *SignerEd25519) ToMgoAddressDevNet() string {
	return s.MgoAddressDevNet
}
func (s *SignerEd25519) ToMgoAddressTestNet() string {
	return s.MgoAddressTestNet
}

func (s *SignerEd25519) Sign(message []byte) []byte {
	return ed25519.Sign(s.PrivateKey, message)
}

func Verify(publicKey []byte, message []byte, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}
func mgoAddressDevNet(publicKey []byte) string {
	tmp := append([]byte{byte(config.Ed25519Flag)}, publicKey...)
	hexHash := utils.Blake2bv1(tmp)
	return "0x" + hex.EncodeToString(hexHash)[:config.AccountAddress32Length*2]
}

func mgoAddressTestNet(publicKey []byte) string {
	inputBytes := append([]byte{byte(config.Ed25519Flag)}, publicKey...)
	return "0x" + hex.EncodeToString(utils.Keccak256(inputBytes))[:config.MGO_ADDRESS_LENGTH]
}
