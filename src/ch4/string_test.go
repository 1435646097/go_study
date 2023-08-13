package ch4

import (
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	str := "我爱中国"
	r := []rune(str)
	log.Println(r)
	//E68891E788B1E4B8ADE59BBD
	log.Printf("%X\n", r)
	log.Println(len(str))
}

func TestStringFunc(t *testing.T) {
	str := "A,C,D,G"
	split := strings.Split(str, ",")
	for _, s := range split {
		t.Log(s)
	}

	join := strings.Join(split, "=")
	t.Log(join)
}

type calculate func(a int) int

func calculationTime(inner calculate) calculate {
	return func(num int) int {
		now := time.Now()
		result := inner(num)
		log.Println("本次函数耗时:", time.Since(now).Milliseconds())
		return result
	}
}

func randomNum(a int) int {
	time.Sleep(1 * time.Second)
	return rand.Intn(a)
}
func TestFunc(t *testing.T) {
	f := calculationTime(randomNum)
	i := f(233)
	t.Log(i)
}
