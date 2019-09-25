package lock_test

import (
	"sync"
	"testing"
	"time"
)

// 设定等待时间版本
func TestLock1(t *testing.T) {
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	t.Log(counter)
}

// 等待
func TestLock2(t *testing.T) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
}
