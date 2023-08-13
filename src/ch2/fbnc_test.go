package ch2

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestFist(t *testing.T) {
	t.Log("hell11o")

}

func TestWori(t *testing.T) {
	t.Log("我的")
}

func TestAsd(t *testing.T) {
	num := 99
	for {
		r := rand.Intn(100) + 1
		if r == num {
			t.Log("恭喜你，猜对了", r)
			break
		}
		if r > num {
			t.Log("你猜的数字太大了", r)
		} else {
			t.Log("你猜的数字小了一丢丢", r)
		}
	}
}

func TestRandomDate(t *testing.T) {
	year := 2018
	rand.Seed(time.Now().Unix())
	month := rand.Intn(12) + 1
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28

	case 4, 6, 7, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1

	fmt.Println(year, month, day)
}

func TestFlotType(t *testing.T) {
	anser := 42.0
	t.Log(reflect.TypeOf(anser))
}

func TestFlotCompute(t *testing.T) {
	f := 0.1
	f += 0.2
	t.Log(f)

	f = 1.0 / 3
	t.Log(f)
	t.Log(f + f + f)
}

func TestPiggyBank(t *testing.T) {
	rand.Seed(time.Now().Unix())
	moneyArray := [...]float64{0.05, 0.10, 0.25}
	var money float64
	for money < 20 {
		index := rand.Intn(len(moneyArray))
		money += moneyArray[index]
		t.Logf("存钱罐里面一共有： %4.2f 美元", money)
	}
}

func TestChar(t *testing.T) {
	// c := "dddd我"
	// t.Log("10" + 1)
	// t.Log(string(65))
}
func showSliceInfo(slice []string) {
	fmt.Printf("len %v ,capacity %v , info %v \n", len(slice), cap(slice), slice)
}

func TestSlice(t *testing.T) {
	array := []string{"1", "2", "3", "4", "5"}
	s := array[2:3:3]
	s[0] = "20"
	s1 := append(s, "555")
	s1 = append(s1, "555")
	s1 = append(s1, "555")
	s1 = append(s1, "555")
	showSliceInfo(array)
	showSliceInfo(s)
	showSliceInfo(s1)
	b, err := json.MarshalIndent(array, "", "")
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(b))
	c := make(chan string)
	fmt.Println(c)
}
