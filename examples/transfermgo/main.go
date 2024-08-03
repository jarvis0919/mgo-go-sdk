package main

import (
	"encoding/json"
	"fmt"
	"mgo-go-sdk/account/signer"
	"mgo-go-sdk/utils"
	"strings"
	"time"
)

func main() {

	// key, err := keypair.FetchKeyPair("AI7lfx9QqrErdT/YxvOrU1472YdajhrRB/1/ryfsnZSw")
	// fmt.Println(key, err)
	// k, err := keypair.DecodeMgoPrivateKey("mgoprivkey1qzu2yj46x4n8p7km586t8qa90l3tgzaw8708sxra5vz9gs2638x65w0sqjd")
	// fmt.Println(k, err)
	// l, err := keypair.EncodeMgoPrivateKey(k.SecretKey, "Ed25519")

	// fmt.Println(l, err)
	var builder strings.Builder
	for i := 0; i < 10; i++ {
		s, err := signer.NewEd25519Signer()
		if err != nil {
			fmt.Println(err)
		}
		SignObject, _ := json.Marshal(struct {
			Address  string `json:"address" form:"address" `
			Time     string `json:"time" form:"time"`
			SignType string `json:"signType" form:"signType"`
		}{
			Address:  s.MgoAddressTestNet,
			Time:     time.Now().Format("2006-01-02"),
			SignType: "Register",
		})
		SignObject2, _ := json.Marshal(struct {
			Address  string `json:"address" form:"address" `
			Time     string `json:"time" form:"time"`
			SignType string `json:"signType" form:"signType"`
		}{
			Address:  s.MgoAddressTestNet,
			Time:     time.Now().Format("2006-01-02"),
			SignType: "Login",
		})
		// fmt.Println("0xee81f95ff8e0d5a4ec13ba1e17fcb0c64720ad589f1acf1bd1884cd3c4e5b309", s.MgoAddressTestNet)
		sign := s.Sign(SignObject)
		data := utils.ByteArrayToBase64String(sign)
		sign2 := s.Sign(SignObject2)
		data2 := utils.ByteArrayToBase64String(sign2)
		// fmt.Println(data)

		builder.WriteString(fmt.Sprintf("PrivateKey: %s\nAddress:    %s\nSignDataRegister:  %s\nSignDataLogin:  %s\n\n", s.PrivateKeyHex(), s.MgoAddressTestNet, data, data2))

	}
	utils.WriteFile("test.txt", builder.String())
	// fmt.Println("address: ", s.MgoAddressDevNet)
	// fmt.Println("PrivateKey: ", s.MgoPrivateKey())
	// fmt.Println("PublicKey: ", s.PublicKeyHex())
	// fmt.Println("Private: ", s.PrivateKeyHex())

	// k, err := signer.NewSignertWithMnemonicDevNet("spring lab bag knock crane rigid now pipe vicious main life job", global.Ed25519Flag)
	// fmt.Println(k, err)

	// fmt.Println("----", s.MgoPrivateKey())
	// // p, err := signer.NewEd25519SignerFromMgoPrivatekey(s.MgoPrivateKey())
	// // fmt.Println(p, err)
	// var ctx = context.Background()
	// var cli = client.NewMgoClient(global.MgoDevnet)
	// rsp, err := cli.TransferMgo(ctx, request.TransferMgoRequest{
	// 	Signer: s,
	// 	// MgoObjectId: "0x9a7291430afe4f989ae45b2469d244a4480433665a6f051a27859c0e5c46027c",
	// 	MgoObjectId: "0x0077e6d729957de6295f4be6c5f673375b1731d15da32a2c3f8238210c63e545",
	// 	GasBudget:   "100000000",
	// 	Recipient:   "0x398c10b4834ea8d243c1ebd63888287cf3ae8a4d0dd320de47be4cba3ab6c6ff",
	// 	Amount:      "1",
	// })
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// utils.JsonPrint(rsp)
	// fmt.Println(base64.StdEncoding.DecodeString(rsp.TxBytes))
	// _, _ := base64.StdEncoding.DecodeString(rsp.TxBytes)
	// fmt.Println(bcs.ULEBDecode(k))
	// rsp2, err := cli.SignAndExecuteTransactionBlock(ctx, request.SignAndExecuteTransactionBlockRequest{
	// 	TxnMetaData: rsp,
	// 	Signer:      s,
	// 	// only fetch the effects field
	// 	Options: request.TransactionBlockOptions{
	// 		ShowInput:    true,
	// 		ShowRawInput: true,
	// 		ShowEffects:  true,
	// 	},
	// 	RequestType: "WaitForLocalExecution",
	// })
	// fmt.Println(err)

	// ok(rsp2)
}

// func Sign(signer signer.Signer, msg string) []byte {
// 	message := []byte(msg)
// 	var header []byte
// 	if len(message) > 128 {
// 		header = []byte{3, 0, 0, byte(len(message)), 1}

// 	} else {
// 		header = []byte{3, 0, 0, byte(len(message))}
// 	}
// 	message = append(header, message...)
// 	message = signer.Sign(Keccak256(message))
// 	public := signer.PublicKeyBytes()
// 	signData := append(message, public...)
// 	signData = append([]byte{0}, signData...)
// 	return signData
// }
// func blake2bv1(message []byte) []byte {
// 	hasher, err := blake2b.New256(nil)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return nil
// 	}
// 	hasher.Write(message)
// 	hash := hasher.Sum(nil)

// 	fmt.Println(hex.EncodeToString(hash))
// 	return hash
// }
// func Keccak256(input []byte) []byte {
// 	hash := sha3.NewLegacyKeccak256()
// 	hash.Write(input)
// 	return hash.Sum(nil)
// }

// func ConvertToTypeSlice(publicKey ed25519.PublicKey) []byte {
// 	typeSlice := make([]byte, len(publicKey))
// 	for i, v := range publicKey {
// 		typeSlice[i] = byte(v)
// 	}
// 	return typeSlice
// }
