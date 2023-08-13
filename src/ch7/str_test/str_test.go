package ch7__test

import (
	"fmt"
	"study/src/ch7/util"
	"testing"
	"unicode/utf8"
)

func TestStrLen(t *testing.T) {
	str := "我爱我的中国a"
	for _, v := range str {
		t.Logf("%c", v)
	}
	r := []rune(str)
	t.Log(len(r))
	t.Log(len(str))
	utf8.RuneCountInString(str)
}

func TestStrUtil(t *testing.T) {
	str := "我爱我的中国a"
	fmt.Printf("util.StrLen(str): %v\n", util.StrLen(str))
}
