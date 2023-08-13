package ch3

import (
	"testing"
)

func TestSlice(t *testing.T) {
	slice1 := []int{1, 2, 3}

	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		t.Log(len(slice1), "---", cap(slice1), "-----", slice1)
	}

	t.Log(slice1)
}

func TestSlice2(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := slice1[1:]

	slice2 = append(slice2, 20)

	slice1[2] = 100
	t.Log(slice1, len(slice1), cap(slice1))
	t.Log(slice2, len(slice2), cap(slice2))
}

func TestMap(t *testing.T) {
	testMap := map[int]int{}
	testMap[1] = 2
	if v, ok := testMap[3]; ok {
		t.Log("这个值确实存在", testMap[1], v)
	}

	t.Log(testMap[1])
}
