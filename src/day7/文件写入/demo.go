package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.OpenFile("1.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开失败")
	}
	defer file.Close()
	//写入
	str := "Golang and Python is my favorite language"
	file.WriteString(str)

}
