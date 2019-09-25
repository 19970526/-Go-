package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("no")
	}
	defer file.Close()
	// 读取
	var tmp = make([]byte, 500)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Println("读取失败")
	}
	if err == io.EOF {
		fmt.Println("没内容")
	}

	fmt.Printf("占%d字节\n", n)
	fmt.Println(tmp[:n])
	fmt.Println("--------------------------------")
	fmt.Println(string(tmp[:n]))

}
