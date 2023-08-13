package main

import (
	"errors"
	"fmt"
)

type Person struct {
	id     int64
	name   string
	gender byte
	age    int
}

func (p *Person) FindOne() {
	prepare, err := db.Prepare("SELECT name,gender,age FROM person WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	r := prepare.QueryRow(p.id)
	err2 := r.Scan(&p.name, &p.gender, &p.age)
	if err2 != nil {
		fmt.Println(err.Error())
	}
}

func (p *Person) Delete() error {
	p.FindOne()
	if len(p.name) == 0 {
		return errors.New("这个id不存在")
	}
	prepare, err := db.Prepare("DELETE FROM person WHERE id = ?")
	if err != nil {
		fmt.Println(err)
	}
	r, err2 := prepare.Exec(p.id)
	if err2 != nil {
		fmt.Println(err.Error())
	}
	result, _ := r.RowsAffected()
	fmt.Println("成功删除了", result, "行")
	return nil
}

func (p *Person) insert() {
	prepare, err := db.Prepare("INSERT INTO `common_mistakes`.`person` (`id`, `name`, `gender`, `age`) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}
	exec, err := prepare.Exec(&p.id, &p.name, &p.gender, &p.age)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("这个主键id是", id)
}
