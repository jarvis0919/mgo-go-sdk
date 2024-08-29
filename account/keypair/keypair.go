package keypair

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/account/signer/ed25519"
	"github.com/jarvis0919/mgo-go-sdk/bcs"
	"github.com/jarvis0919/mgo-go-sdk/config"
	"github.com/jarvis0919/mgo-go-sdk/model"
	"github.com/jarvis0919/mgo-go-sdk/utils"
)

type Options struct {
	Scheme     config.Scheme
	PrivateKey string
}

type Keypair struct {
	signer.Signer
	Scheme config.Scheme
}

func New(opt Options) (*Keypair, error) {
	switch opt.Scheme {
	case config.Secp256k1Flag:
		return nil, errors.New("invalid signature scheme flag")
	case config.Ed25519Flag:
		if opt.PrivateKey == "" {
			sig, err := ed25519.NewEd25519Signer()
			if err != nil {
				return nil, err
			}
			return &Keypair{Scheme: opt.Scheme, Signer: sig}, err
		}
		if regexp.MustCompile(`^(0x|0X)?[0-9a-fA-F]+$`).MatchString(opt.PrivateKey) {
			privateKey := opt.PrivateKey
			if strings.HasPrefix(privateKey, "0x") || strings.HasPrefix(privateKey, "0X") {
				privateKey = privateKey[2:]
			}
			if len(privateKey) != 64 {
				return nil, errors.New("invalid private key")
			}
			sig, err := ed25519.NewEd25519SignerFromPrivateKey(privateKey)
			if err != nil {
				return nil, err
			}
			return &Keypair{Scheme: opt.Scheme, Signer: sig}, err
		}
		sig, err := ed25519.NewEd25519SignerFromMgoPrivatekey(opt.PrivateKey)
		if err != nil {
			return nil, err
		}
		return &Keypair{Scheme: opt.Scheme, Signer: sig}, err
	default:
		return nil, errors.New("invalid signature scheme flag")
	}
}

type SignedTransactionSerializedSig struct {
	TxBytes   string `json:"tx_bytes"  yaml:"txBytes"`
	Signature string `json:"signature" yaml:"signature"`
}

func (k *Keypair) SignPersonalMessage(message []byte, net config.NetIdentity) []byte {
	message = append(bcs.ULEBEncode(uint64(len(message))), message...)
	data := dataWithIntent(message, config.PersonalMessage)
	digest := digestData(data, net)
	sigBytes := k.Sign(digest[:])
	publicKey := k.PublicKeyBytes()

	signData := append(sigBytes, publicKey...)
	signData = append([]byte{byte(config.Ed25519Flag)}, signData...)
	return signData
}

func (k *Keypair) SignTransactionBlock(txn *model.TxnMetaData, net config.NetIdentity) *SignedTransactionSerializedSig {
	txBytes, _ := base64.StdEncoding.DecodeString(txn.TxBytes)
	data := dataWithIntent(txBytes, config.TransactionData)
	digest := digestData(data, net)

	sigBytes := k.Sign(digest[:])

	return &SignedTransactionSerializedSig{
		TxBytes:   txn.TxBytes,
		Signature: k.toSerializedSignature(sigBytes),
	}
}

func dataWithIntent(data []byte, intent config.Keytype) []byte {
	header := []byte{byte(intent), 0, 0}
	markData := make([]byte, len(header)+len(data))
	copy(markData, header)
	copy(markData[len(header):], data)
	return markData
}
func digestData(data []byte, net config.NetIdentity) []byte {
	if net == config.MgoTestnet {
		return utils.Keccak256(data)
	} else {
		return utils.Blake2bv1(data)
	}

}
func (k *Keypair) toSerializedSignature(signature []byte) string {
	signatureLen := len(signature)
	pubKeyLen := len(k.PublicKeyBytes())
	serializedSignature := make([]byte, 1+signatureLen+pubKeyLen)
	serializedSignature[0] = byte(k.Scheme)
	copy(serializedSignature[1:], signature)
	copy(serializedSignature[1+signatureLen:], k.PublicKeyBytes())
	return base64.StdEncoding.EncodeToString(serializedSignature)
}

type SignatureInfo struct {
	SerializedSignature []byte
	SignatureScheme     string
	Signature           []byte
	PublicKey           []byte
	Bytes               []byte
}

func VerifyPersonalMessage(msg []byte, sig []byte, net config.NetIdentity) bool {
	scheme := getSignatureScheme(sig)
	if sig == nil || scheme == "" || len(sig) != config.SIGNATURE_SCHEME_TO_SIZE[scheme] {
		return false
	}
	signatureInfo := parseSignatureInfo(sig, scheme)
	publickey := signatureInfo.PublicKey
	msgReserialize := append(bcs.ULEBEncode(uint64(len(msg))), msg...)
	intentMessage := dataWithIntent(msgReserialize, config.PersonalMessage)
	digest := digestData(intentMessage, net)

	switch config.SIGNATURE_SCHEME_TO_FLAG[scheme] {
	case config.Ed25519Flag:
		return ed25519.Verify(publickey, digest, signatureInfo.Signature)
	// case config.Secp256k1Flag:
	// 	secp256k1.Verify(publickey, digest, signatureInfo.Signature)
	default:
		return false
	}
}

func VerifyTransactionBlock(txn []byte, sig []byte, net config.NetIdentity) bool {
	scheme := getSignatureScheme(sig)
	if sig == nil || scheme == "" {
		return false
	}
	signatureInfo := parseSignatureInfo(sig, scheme)
	publickey := signatureInfo.PublicKey

	intentMessage := dataWithIntent(txn, config.TransactionData)
	digest := digestData(intentMessage, net)

	switch config.SIGNATURE_SCHEME_TO_FLAG[scheme] {
	case config.Ed25519Flag:
		return ed25519.Verify(publickey, digest, signatureInfo.Signature)
	default:
		return false
	}
}
func getSignatureScheme(bytes []byte) string {
	return config.SIGNATURE_FLAG_TO_SCHEME[config.Scheme(bytes[0])]
}

func parseSignatureInfo(bytes []byte, signatureScheme string) *SignatureInfo {
	size := config.SIGNATURE_SCHEME_TO_SIZE[signatureScheme]
	signature := bytes[1 : len(bytes)-size]
	publicKey := bytes[1+len(signature):]

	return &SignatureInfo{
		SerializedSignature: bytes,
		SignatureScheme:     signatureScheme,
		Signature:           signature,
		PublicKey:           publicKey,
		Bytes:               bytes,
	}
}
