package ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/global"

	// "github.com/jarvis0919/mgo-go-sdk/keypair"

	"github.com/jarvis0919/mgo-go-sdk/utils"
)

// SignerEd25519 结构体定义
type SignerEd25519 struct {
	PrivateKey        ed25519.PrivateKey // 私钥 secretKey
	PublicKey         ed25519.PublicKey  // 公钥 peerId
	MgoAddressDevNet  string             // 地址
	MgoAddressTestNet string
}

// NewEd25519Signer 创建新的 Ed25519 签名器
func NewEd25519Signer() (*SignerEd25519, error) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return newEd25519Signer(privateKey)
}
func NewEd25519SignerFromPrivateKey(mgoPrivatekey string) (ed25519Signer *SignerEd25519, err error) {
	seed, err := hex.DecodeString(mgoPrivatekey)
	if err != nil {
		return nil, err
	}
	return newEd25519SignerFromSeed(seed)
}

// NewEd25519SignerFromMgoPrivatekey 从 Mgo 私钥创建新的 Ed25519 签名器
func NewEd25519SignerFromMgoPrivatekey(mgoPrivatekey string) (ed25519Signer *SignerEd25519, err error) {
	ParsedKeypair, err := signer.DecodeMgoPrivateKey(mgoPrivatekey)
	if err != nil {
		return nil, err
	}
	return newEd25519SignerFromSeed(ParsedKeypair.SecretKey)
}

// newEd25519SignerFromSeed 从种子创建新的 Ed25519 签名器
func newEd25519SignerFromSeed(seed []byte) (ed25519Signer *SignerEd25519, err error) {
	defer func() {
		if r := recover(); r != nil {
			ed25519Signer = nil
			err = fmt.Errorf("recovered from panic: %s", r)
		}
	}()
	return newEd25519Signer(ed25519.NewKeyFromSeed(seed))
}

// newEd25519Signer 初始化 SignerEd25519 结构体
func newEd25519Signer(privateKey ed25519.PrivateKey) (*SignerEd25519, error) {
	publicKey := privateKey.Public().(ed25519.PublicKey)
	// flag, ok := global.SIGNATURE_SCHEME_TO_FLAG["Ed25519"]
	// if !ok {
	// 	return nil, errors.New("invalid signature scheme flag")
	// }

	return &SignerEd25519{
		PrivateKey:        privateKey,
		PublicKey:         publicKey,
		MgoAddressDevNet:  mgoAddressDevNet(publicKey),
		MgoAddressTestNet: mgoAddressTestNet(publicKey),
	}, nil
}

// Sign 签名消息
func (s *SignerEd25519) Sign(message []byte) []byte {
	// var header []byte
	// header = []byte{byte(global.PersonalMessage), 0, 0}
	// header = append(header, bcs.ULEBEncode(uint64(len(message)))...)
	// message = append(header, message...)
	// message = ed25519.Sign(s.PrivateKey, utils.Blake2bv1(message))
	// public := s.PublicKeyBytes()
	// signData := append(message, public...)
	// signData = append([]byte{byte(global.Ed25519Flag)}, signData...)
	return ed25519.Sign(s.PrivateKey, message)
}

// String 返回 SignerEd25519 的字符串表示
func (s SignerEd25519) String() string {
	return fmt.Sprintf("PrivateKey: %s\nPublicKey: %s\nMgoAddressDevNet: %s\nMgoAddressTestNet: %s",
		hex.EncodeToString(s.PrivateKey),
		hex.EncodeToString(s.PublicKey),
		s.MgoAddressDevNet, s.MgoAddressTestNet)
}

// // MgoPrivateKey 返回 Mgo 私钥
// func (s *SignerEd25519) MgoPrivateKey() string {
// 	mgoPrivateKey, _ := signer.EncodeMgoPrivateKey(s.PrivateKey[:global.PRIVATE_KEY_SIZE], s.Scheme)
// 	return mgoPrivateKey
// }

// SecretKeyHex 返回十六进制表示的私钥
func (s *SignerEd25519) SecretKeyHex() string {
	return "0x" + hex.EncodeToString(s.PrivateKey)
}

// SecretKeyBytes 返回字节数组表示的私钥
func (s *SignerEd25519) SecretKeyBytes() []byte {
	return s.PrivateKey
}

// PrivateKeyHex 返回十六进制表示的私钥
func (s *SignerEd25519) PrivateKeyHex() string {
	return "0x" + hex.EncodeToString(s.PrivateKey[:32])
}

// PrivateKeyBytes 返回字节数组表示的私钥
func (s *SignerEd25519) PrivateKeyBytes() []byte {
	return s.PrivateKey[:32]
}

// PublicKeyHex 返回十六进制表示的公钥
func (s *SignerEd25519) PublicKeyHex() string {
	return "0x" + hex.EncodeToString(s.PublicKey)
}

// PublicKeyBytes 返回字节数组表示的公钥
func (s *SignerEd25519) PublicKeyBytes() []byte {
	return s.PublicKey
}

// PublicBase64Key 返回 Base64 编码的公钥
func (s *SignerEd25519) PublicBase64Key() string {
	return utils.EncodeBase64(s.PublicKey)
}
func (s *SignerEd25519) ToMgoAddressDevNet() string {
	return s.MgoAddressDevNet
}
func (s *SignerEd25519) ToMgoAddressTestNet() string {
	return s.MgoAddressTestNet
}

// mgoAddressDevNet 返回开发网络中的 Mgo 地址
func mgoAddressDevNet(publicKey []byte) string {
	tmp := append([]byte{byte(global.Ed25519Flag)}, publicKey...)
	hexHash := utils.Blake2bv1(tmp)
	return "0x" + hex.EncodeToString(hexHash)[:global.AccountAddress32Length*2]
}

func mgoAddressTestNet(publicKey []byte) string {
	inputBytes := append([]byte{byte(global.Ed25519Flag)}, publicKey...)
	return "0x" + hex.EncodeToString(utils.Keccak256(inputBytes))[:global.MGO_ADDRESS_LENGTH]
}
