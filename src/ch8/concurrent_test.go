package ch8

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 500)
}

func TestUnsafe(t *testing.T) {
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			count++
		}()
	}

	time.Sleep(time.Millisecond * 1500)

	t.Log(count)
}

func TestSafe(t *testing.T) {
	count := 0
	lock := sync.Mutex{}
	for i := 0; i < 5000; i++ {
		go func() {
			lock.Lock()
			defer func() {
				lock.Unlock()
			}()
			count++
		}()

	}
	time.Sleep(time.Millisecond * 600)
	t.Log(count)
}

func TestWaitGroup(t *testing.T) {
	count := 0
	waitGroup := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 0; i < 5000; i++ {
		waitGroup.Add(1)
		go func() {
			lock.Lock()
			defer func() {
				lock.Unlock()
				waitGroup.Done()
			}()
			count++
		}()
	}
	waitGroup.Wait()
	t.Log(count)
}
