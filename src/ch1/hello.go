package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) string() {
	p.Name = "我的你的码"
}

func main() {
	day := 28
	distance := 56000000
	speed := distance / day
	fmt.Println("这个平均速度为必须为：", speed, "公路/小时")
}
