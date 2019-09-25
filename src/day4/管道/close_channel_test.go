package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

//写进管道
func ComeIn(c chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//关闭管道
		close(c)
		wg.Done()
	}()

}

//从管道读取
func ComeOut(c chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChannelComeAndOut(t *testing.T) {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(1)
	ComeIn(c, &wg)
	wg.Add(1)
	ComeOut(c, &wg)
	wg.Wait()
}
