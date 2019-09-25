package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func InChannel() string {
	time.Sleep(time.Millisecond * 500)
	return "this is a channel"
}

func WriteChannel() chan string {
	c := make(chan string, 1)
	go func() {
		ret := InChannel()
		fmt.Println("start channel")
		c <- ret
		fmt.Println("end channel")
	}()
	return c
}

func TestChannel(t *testing.T) {
	select {
	case result := <-WriteChannel():
		t.Log(result)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
