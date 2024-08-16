package keypair

import (
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/account/signer/ed25519"
	"github.com/jarvis0919/mgo-go-sdk/bcs"
	"github.com/jarvis0919/mgo-go-sdk/global"
	"github.com/jarvis0919/mgo-go-sdk/model"
	"github.com/jarvis0919/mgo-go-sdk/utils"
)

type Options struct {
	Scheme     global.Scheme
	PrivateKey string
}

type Keypair struct {
	signer.Signer
	Scheme global.Scheme
}

func New(opt Options) (*Keypair, error) {
	switch opt.Scheme {
	case global.Secp256k1Flag:
		return nil, errors.New("invalid signature scheme flag")
	case global.Ed25519Flag:
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

func (k *Keypair) SignPersonalMessage(message []byte, net global.NetIdentity) []byte {
	message = append(bcs.ULEBEncode(uint64(len(message))), message...)
	data := k.dataWithIntent(message, global.PersonalMessage)
	digest := k.digestData(data, net)
	sigBytes := k.Sign(digest[:])
	publicKey := k.PublicKeyBytes()

	signData := append(sigBytes, publicKey...)
	signData = append([]byte{byte(global.Ed25519Flag)}, signData...)
	return signData
}

func (k *Keypair) SignTransactionBlock(txn *model.TxnMetaData, net global.NetIdentity) *SignedTransactionSerializedSig {
	txBytes, _ := base64.StdEncoding.DecodeString(txn.TxBytes)
	data := k.dataWithIntent(txBytes, global.TransactionData)
	digest := k.digestData(data, net)

	sigBytes := k.Sign(digest[:])

	return &SignedTransactionSerializedSig{
		TxBytes:   txn.TxBytes,
		Signature: k.toSerializedSignature(sigBytes),
	}
}

func (k *Keypair) dataWithIntent(data []byte, intent global.Keytype) []byte {
	header := []byte{byte(intent), 0, 0}
	markData := make([]byte, len(header)+len(data))
	copy(markData, header)
	copy(markData[len(header):], data)
	return markData
}
func (k *Keypair) digestData(data []byte, net global.NetIdentity) []byte {
	if net == global.MgoTestnet {
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
