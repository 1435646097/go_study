package ch5

import (
	"fmt"
	"testing"
)

type List struct {
	size int
}

func (list *List) add() {
	list.size++

}

type ArrayList struct {
	List
}

func (list *ArrayList) add() {
	list.size += 2
}

func TestExtend(t *testing.T) {
	slice1 := []int{1, 2, 3}
	arr := []interface{}{3, 4, 5, true}
	list := ArrayList{}
	list.add()
	list.add()
	list.add()
	fmt.Printf("list.size: %v\n", list.size)
	fmt.Println(arr...)
	fmt.Println(slice1)
}
