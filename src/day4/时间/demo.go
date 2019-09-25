package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	date := time.Now()
	fmt.Println(date)
	// add
	fmt.Println(date.Add(time.Hour))
	// sub
	date2 := date.AddDate(0, 0, 3)
	fmt.Println(date.Sub(date2))
	//格式化
	fmt.Println(date.Format("2006-01-02"))
	fmt.Println(date.Format("2006-01-02 15:04:05"))
	fmt.Println(date.Format("2006/1/2"))
	fmt.Println(date.Format("2006/1/2 15:04:05"))
	// Go语言的发明时间是2006年1月2日15点零4分零5秒，简单记成2006,1,2,3,4,5
}
