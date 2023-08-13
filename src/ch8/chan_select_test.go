package ch8

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type Person struct{}

func TaskJob1() chan bool {
	c := make(chan bool, 1)
	time.Sleep(500 * time.Millisecond)
	c <- true
	return c
}

func TestChan1(t *testing.T) {
	select {
	case c := <-FutureTask():
		t.Log("成功取出来了值", c)
	case d := <-time.After(5 * time.Second):
		t.Error("取不出来咯", d)
	}
	t.Log("操")
	p := new(Person)
	fmt.Printf("unsafe.Pointer(p): %v\n", unsafe.Pointer(p))
}
