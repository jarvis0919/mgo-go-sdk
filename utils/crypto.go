package utils

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func Keccak256(input []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(input)
	return hash.Sum(nil)
}
func Blake2bv1(message []byte) []byte {
	hasher, err := blake2b.New256(nil)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	hasher.Write(message)
	hash := hasher.Sum(nil)
	return hash
}
func EncodeBase64(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

func DecodeBase64(value string) []byte {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil
	}
	return data
}

// HexStringToByteArray 将16进制字符串转换为字节数组
func HexStringToByteArray(hexString string) ([]byte, error) {
	return hex.DecodeString(hexString)
}

// ByteArrayToHexString 将字节数组转换为16进制字符串
func ByteArrayToHexString(byteArray []byte) string {
	return hex.EncodeToString(byteArray)
}

func ByteArrayToBase64String(byteArray []byte) string {
	return base64.StdEncoding.EncodeToString(byteArray)
}

func Base64StringToByteArray(base64String string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64String)
}
func JsonPrint(v any) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("JSON 编码失败: %s", err)
	}
	// 将JSON字节数据转换为字符串并打印
	fmt.Println(string(jsonData))
}
