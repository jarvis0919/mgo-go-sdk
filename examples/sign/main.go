package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jarvis0919/mgo-go-sdk/account/signer"
	"github.com/jarvis0919/mgo-go-sdk/utils"

	"github.com/ethereum/go-ethereum/crypto"
)

type SignObject struct {
	Address  string `json:"address"  yaml:"address"  form:"address"`
	Time     string `json:"time"     yaml:"time"     form:"time"`
	SignType string `json:"signType" yaml:"signType" form:"signType"`
}

func main() {

}

// // ERC20ABI is the ABI for the ERC-20 token contract (simplified version).
// const ERC20ABI = `[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"}]`

// func SendTransaction(receiver string, amount *big.Int) (string, error) {
// 	// 连接到以太坊节点
// 	client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
// 	if err != nil {
// 		return "", fmt.Errorf("failed to connect to the Ethereum client: %w", err)
// 	}

// 	// 发起交易者的私钥
// 	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY_HERE")
// 	if err != nil {
// 		return "", fmt.Errorf("failed to load private key: %w", err)
// 	}

// 	// 获取发起交易者的地址
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		return "", errors.New("error casting public key to ECDSA")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

// 	// 获取交易计数
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get the nonce: %w", err)
// 	}

// 	// 获取建议的Gas价格
// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return "", fmt.Errorf("failed to suggest gas price: %w", err)
// 	}

// 	// 获取链ID
// 	chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		return "", fmt.Errorf("failed to get chain ID: %w", err)
// 	}

// 	// 创建交易认证
// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create authorized transactor: %w", err)
// 	}
// 	auth.From = fromAddress
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // 发送的ETH数量
// 	auth.GasLimit = uint64(300000) // 设置Gas限制
// 	auth.GasPrice = gasPrice

// 	// ERC-20 代币合约地址
// 	tokenAddress := common.HexToAddress("0x8FE1Fc07eA6928aae8c6bB90A5CC5A992407932F")

// 	// 接收者地址
// 	toAddress := common.HexToAddress(receiver)

// 	// 调用ERC-20合约的转账方法
// 	transferFnSignature := []byte("transfer(address,uint256)")
// 	hash := crypto.Keccak256Hash(transferFnSignature)
// 	methodID := hash[:4]

// 	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
// 	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

// 	var data []byte
// 	data = append(data, methodID...)
// 	data = append(data, paddedAddress...)
// 	data = append(data, paddedAmount...)

// 	// 估算Gas限制
// 	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
// 		From: fromAddress,
// 		To:   &tokenAddress,
// 		Data: data,
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("failed to estimate gas limit: %w", err)
// 	}

// 	// 创建并签署交易
// 	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)
// 	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to sign transaction: %w", err)
// 	}

// 	// 发送交易
// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to send transaction: %w", err)
// 	}
// 	return signedTx.Hash().Hex(), nil
// }

func Secp256k1Sign() {

	privateKeyHex := "69292f54d754a1c1c7bba52966ecd45b3b87e22c45c33eb1319066a322c40747" // 使用你的私钥
	SignObject, err := json.Marshal(SignObject{
		Address:  "0x8ae9Ae6b71e14023fb3C81A6fe5d30aBe771Cc1E",
		Time:     time.Now().Format("2006-01-02"),
		SignType: "sdLogin",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(SignObject))
	fmt.Println(utils.ByteArrayToHexString(utils.Keccak256(SignObject)))
	preamble := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(SignObject)))
	ethMessage := append(preamble, []byte(SignObject)...)
	// 计算 Keccak-256 哈希值
	hash := utils.Keccak256(ethMessage)
	fmt.Printf("Keccak-256 哈希值: %x\n", hash)

	prv, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" 公钥:", utils.ByteArrayToHexString(crypto.CompressPubkey(&prv.PublicKey)))
	// bytekey, _ := utils.HexStringToByteArray(utils.ByteArrayToHexString())

	fmt.Println(utils.ByteArrayToHexString(crypto.FromECDSA(prv)))
	sig, err := crypto.Sign(hash, prv)

	fmt.Println(utils.ByteArrayToHexString(sig), err)
	fmt.Println(crypto.VerifySignature(crypto.CompressPubkey(&prv.PublicKey), hash, sig[:64]))
	fmt.Println(utils.ByteArrayToHexString(sig[:64]), err)
}

func Ed25519Sign() {

	s, err := signer.NewEd25519Signer()
	if err != nil {
		fmt.Println(err)
	}
	SignObject, _ := json.Marshal(struct {
		Address  string `json:"address"  yaml:"address"  form:"address"`
		Time     string `json:"time"     yaml:"time"     form:"time"`
		SignType string `json:"signType" yaml:"signType" form:"signType"`
	}{
		Address:  s.MgoAddressTestNet,
		Time:     time.Now().Format("2006-01-02"),
		SignType: "Register",
	})
	// SignObject2, _ := json.Marshal(struct {
	// 	Address  string `json:"address"  yaml:"address"  form:"address"`
	// 	Time     string `json:"time"     yaml:"time"     form:"time"`
	// 	SignType string `json:"signType" yaml:"signType" form:"signType"`
	// }{
	// 	Address:  s.MgoAddressTestNet,
	// 	Time:     time.Now().Format("2006-01-02"),
	// 	SignType: "Login",
	// })
	// fmt.Println("0xee81f95ff8e0d5a4ec13ba1e17fcb0c64720ad589f1acf1bd1884cd3c4e5b309", s.MgoAddressTestNet)
	sign := s.Sign(SignObject)
	data := utils.ByteArrayToBase64String(sign)
	fmt.Println(data)
	// sign2 := s.Sign(SignObject2)
	// data2 := utils.ByteArrayToBase64String(sign2)
	// var builder strings.Builder
	// for i := 0; i < 10; i++ {
	// s, err := crypto.GenerateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// SignObject, _ := json.Marshal(struct {
	// 	Address  string `json:"address"  yaml:"address"  form:"address"`
	// 	Time     string `json:"time"     yaml:"time"     form:"time"`
	// 	SignType string `json:"signType" yaml:"signType" form:"signType"`
	// }{
	// 	Address:  utils.ByteArrayToHexString(crypto.CompressPubkey(&s.PublicKey)),
	// 	Time:     time.Now().Format("2006-01-02"),
	// 	SignType: "sdRegister",
	// })
	// 	SignObject2, _ := json.Marshal(struct {
	// 		Address  string `json:"address"  yaml:"address"  form:"address"`
	// 		Time     string `json:"time"     yaml:"time"     form:"time"`
	// 		SignType string `json:"signType" yaml:"signType" form:"signType"`
	// 	}{
	// 		Address:  utils.ByteArrayToHexString(crypto.CompressPubkey(&s.PublicKey)),
	// 		Time:     time.Now().Format("2006-01-02"),
	// 		SignType: "sdLogin",
	// 	})
	// 	// fmt.Println("0xee81f95ff8e0d5a4ec13ba1e17fcb0c64720ad589f1acf1bd1884cd3c4e5b309", s.MgoAddressTestNet)
	// preamble := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(SignObject)))
	// ethMessage := append(preamble, []byte(SignObject)...)
	// // 计算 Keccak-256 哈希值
	// hash := utils.Keccak256(ethMessage)
	// 	preamble2 := []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(SignObject2)))
	// 	ethMessage2 := append(preamble2, []byte(SignObject2)...)
	// 	// 计算 Keccak-256 哈希值
	// 	hash2 := utils.Keccak256(ethMessage2)
	// 	sign, _ := crypto.Sign(hash, s)

	// 	sign2, _ := crypto.Sign(hash2, s)

	// 	// fmt.Println(data)

	// 	builder.WriteString(fmt.Sprintf("PrivateKey: %s\nAddress:    %s\nSignDataRegister:  %s\nSignDataLogin:  %s\n\n", utils.ByteArrayToHexString(crypto.FromECDSA(s)), utils.ByteArrayToHexString(crypto.CompressPubkey(&s.PublicKey)), utils.ByteArrayToHexString(sign), utils.ByteArrayToHexString(sign2)))

	// }
	// utils.WriteFile("test.txt", builder.String())

	// v := sig[0] - 27
	// sigPublicKey, err := crypto.SigToPub(hash, sig)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(utils.ByteArrayToHexString(crypto.CompressPubkey(&prv.PublicKey)))
	// fmt.Println(utils.ByteArrayToHexString(crypto.CompressPubkey(sigPublicKey)))
	// 79dfe960368d341a5dec8df1592cd206985f3f6aa59160c81e001e8e8ee46877226d2f76c8c42a06c41879314b7975c6a18add900ee76766255c4c5e9e578c501b
	// 79dfe960368d341a5dec8df1592cd206985f3f6aa59160c81e001e8e8ee46877226d2f76c8c42a06c41879314b7975c6a18add900ee76766255c4c5e9e578c5000

	// fmt.Println(SendTransaction("0x2330fcFB012c740F7100570Fe5B5bba8f6704482", big.NewInt(1000000000000000000)))
}
