package main

import (
	"fmt"
	//"os"
	"flag"
)

var (
	name string
	age  int
)

func main() {
	// 常规操作用os包
	// for i, j := range os.Args {
	// 	fmt.Printf("args[%d] is %v\n", i, j)
	// }
	flag.StringVar(&name, "name", "同学", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.Parse()
	fmt.Printf("姓名：%s，年龄：%d\n", name, age)
}
