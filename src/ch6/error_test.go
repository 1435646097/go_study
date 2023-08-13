package ch6

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err1 := errors.New("hello")
	err2 := errors.New("hello")
	if err1.Error() == err2.Error() {
		t.Log("确实相等")
		return
	}
	t.Log("不相等的小老弟")
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("错误出现了,", err)
		}
	}()

	t.Log("11111")

	panic(errors.New("来了"))
}
