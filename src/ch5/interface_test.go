package ch5

import (
	"fmt"
	"testing"
)

type Program interface {
	Say()
	Write()
}

type GoProgram struct {
}

func (p *GoProgram) Say() {
	fmt.Println("我是Golang")
}
func (p *GoProgram) Write() {
	fmt.Println("不好意思,go我也会写")
}

type JavaProgram struct {
}

func (p *JavaProgram) Say() {
	fmt.Println("我是JAVA")
}

func (p *JavaProgram) Write() {
	fmt.Println("我写的是JAVA")
}

func TestInteface(t *testing.T) {
	var program Program = new(GoProgram)
	program.Write()

}
