package keypair

import (
	"encoding/base64"
	"encoding/hex"
	"mgo-go-sdk/global"
	"mgo-go-sdk/model"
	"mgo-go-sdk/utils"
)

func FetchKeyPair(value string) (model.MgoKeyPair, error) {
	result, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return model.MgoKeyPair{}, err
	}
	if len(result) == 0 {
		return model.MgoKeyPair{}, err
	}
	switch result[0] {
	case byte(global.Ed25519Flag):
		pb := result[1 : global.Ed25519PublicKeyLength+1]
		sk := result[1+global.Ed25519PublicKeyLength:]
		pbInBase64 := utils.EncodeBase64(pb)
		return model.MgoKeyPair{
			Flag:            byte(global.Ed25519Flag),
			PrivateKey:      sk,
			PublicKeyBase64: pbInBase64,
			PublicKey:       pb,
			MgoAddress:      fromPublicKeyBytesToAddress(pb, byte(global.Ed25519Flag)),
		}, nil
	case byte(global.Secp256k1Flag):
		pb := result[1 : global.Secp256k1PublicKeyLength+1]
		sk := result[1+global.Secp256k1PublicKeyLength:]
		pbInBase64 := utils.EncodeBase64(pb)
		return model.MgoKeyPair{
			Flag:            byte(global.Secp256k1Flag),
			PrivateKey:      sk,
			PublicKey:       pb,
			PublicKeyBase64: pbInBase64,
			MgoAddress:      fromPublicKeyBytesToAddress(pb, byte(global.Secp256k1Flag)),
		}, nil
	default:
		return model.MgoKeyPair{}, global.ErrInvalidEncryptFlag
	}
}

func fromPublicKeyBytesToAddress(publicKey []byte, scheme byte) string {
	if scheme != byte(global.Ed25519Flag) && scheme != byte(global.Secp256k1Flag) {
		return ""
	}
	// 注释使用方法性能更高,但为了维护性，采用了append的方式处理
	// tmp := make([]byte, len(publicKey)+1)
	// tmp[0] = scheme
	// for i := range publicKey {
	// 	tmp[i+1] = publicKey[i]
	// }
	tmp := append([]byte{scheme}, publicKey...)
	hexHash := utils.Blake2bv1(tmp)
	return "0x" + hex.EncodeToString(hexHash[:])[:global.AccountAddress32Length*2]
}
