package ch10

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func (p *Person) SayHello(name string) {
	fmt.Println("你好啊", name)
}

func TestReflectInfo(t *testing.T) {
	p := new(Person)
	// p.Name = "我的名字"
	// p.Age = 30
	t.Log("Type Of", reflect.TypeOf(p))
	t.Log("Value Of", reflect.ValueOf(p))

	filed := reflect.ValueOf(p).Elem().FieldByName("Name")
	// t.Log(reflect.ValueOf(filed))
	// t.Log(filed.Type().Name())
	t.Log("字段的值为", filed)
	t.Log(p)

	reflect.ValueOf(p).MethodByName("SayHello").Call([]reflect.Value{reflect.ValueOf("小张")})
	m := sync.Mutex{}
	m.Unlock()
}
