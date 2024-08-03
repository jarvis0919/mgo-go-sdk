package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type GETHData struct {
	Jsonrpc string `json:"jsonrpc" yaml:"jsonrpc"`
	ID      int    `json:"id"      yaml:"iD"`
	Result  Block  `json:"result"  yaml:"result"`
}
type Block struct {
	BlockHash            string        `json:"blockHash"            yaml:"blockHash"`
	BlockNumber          string        `json:"blockNumber"          yaml:"blockNumber"`
	From                 string        `json:"from"                 yaml:"from"`
	Gas                  string        `json:"gas"                  yaml:"gas"`
	GasPrice             string        `json:"gasPrice"             yaml:"gasPrice"`
	MaxFeePerGas         string        `json:"maxFeePerGas"         yaml:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas" yaml:"maxPriorityFeePerGas"`
	Hash                 string        `json:"hash"                 yaml:"hash"`
	Input                string        `json:"input"                yaml:"input"`
	Nonce                string        `json:"nonce"                yaml:"nonce"`
	To                   string        `json:"to"                   yaml:"to"`
	TransactionIndex     string        `json:"transactionIndex"     yaml:"transactionIndex"`
	Value                string        `json:"value"                yaml:"value"`
	Type                 string        `json:"type"                 yaml:"type"`
	AccessList           []interface{} `json:"accessList"           yaml:"accessList"`
	ChainID              string        `json:"chainId"              yaml:"chainID"`
	V                    string        `json:"v"                    yaml:"v"`
	R                    string        `json:"r"                    yaml:"r"`
	S                    string        `json:"s"                    yaml:"s"`
	YParity              string        `json:"yParity"              yaml:"yParity"`
}

func main() {
	client := resty.New()
	baseURL := "https://api-testnet.bscscan.com/api"
	module := "proxy"
	action := "eth_getTransactionByHash"
	txhash := "0xf679470e0416c57128b9407992cbfddfaec4cf9fd7041dc3ec478cd93721545c"
	apikey := "7K31VQWR41RQNTSWQJ8BA1R2CJJH5NEDG3"

	var blockData GETHData
	_, err := client.R().
		SetQueryParams(map[string]string{
			"module": module,
			"action": action,
			"txhash": txhash,
			"apikey": apikey,
		}).SetResult(&blockData).
		Get(baseURL)
	if err != nil {
		// global.GVA_LOG.Error("Error while making request:", zap.Error(err))
		return
	}

	if blockData.Result.BlockNumber == "" {
		// global.GVA_LOG.Error("", zap.Error(err))
		return
	}

	fmt.Println(blockData.Result.Input[2:10])
	fmt.Println(blockData.Result.Input[10:74])
	fmt.Println(blockData.Result.Input[74:])
	// Calculate 10^18
	// exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	// expRat := new(big.Rat).SetInt(exp)
	// fmt.Println("10^18 as a rational:", expRat)

	// oneMinusDrawFee := big.NewInt(int64((2.1333333318 - 0.1) * 1e10))

	// // Create a rational number representing 0.9 (9/10)
	// ten := big.NewInt(1e10)
	// factor := new(big.Rat).SetFrac(oneMinusDrawFee, ten)

	// // Calculate 0.9 * 10^18
	// resultRat := new(big.Rat).Mul(factor, expRat)

	// resultInt := new(big.Int)
	// resultRat.Num().Div(resultRat.Num(), resultRat.Denom())
	// resultInt = resultRat.Num()
	// fmt.Println("Hello, World! as an integer:", resultInt)
}
