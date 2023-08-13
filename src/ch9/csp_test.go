package ch9

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		go once.Do(func() {
			t.Log("你好世界")
		})
	}
	time.Sleep(500 * time.Millisecond)
}

func TestAny(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	c := make(chan string, 10)

	for i := 0; i < 10; i++ {
		go func() {
			c <- "hello"
		}()
	}
	<-c
	t.Log("完成了")
	time.Sleep(500 * time.Millisecond)
	t.Log(runtime.NumGoroutine())
}
