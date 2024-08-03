package utils

import (
	"fmt"
	"os"
)

func WriteFile(fileName string, builder string) {

	// 打开文件用于追加内容，如果文件不存在则创建它
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// 写入文件
	_, err = file.WriteString(builder)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		file.Close()
		return
	}
	// 关闭文件
	file.Close()

	// fmt.Printf("Batch of %d key pairs written to %s\n", batchSize, fileName)
}
