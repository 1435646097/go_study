package main

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/common_mistakes")
	if err != nil {
		fmt.Println(err)
	}
}

func findById(id int) (*Person, error) {
	p := &Person{}
	s, err := db.Prepare("SELECT id,name,gender,age FROM person where id = ?")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer s.Close()
	r := s.QueryRow(1542141408138334210)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// 	return nil, err2
	// }
	r.Scan(&p.id, &p.name, &p.gender, &p.age)
	return p, err
}

func findAll() (result []Person) {
	prepare, err := db.Prepare("SELECT id,name,gender,age FROM person")
	if err != nil {
		fmt.Println(err.Error())
	}
	r, err2 := prepare.Query()
	if err2 != nil {
		fmt.Println(err.Error())
	}
	for r.Next() {
		p := Person{}
		r.Scan(&p.id, &p.name, &p.gender, &p.age)
		result = append(result, p)
	}
	return result
}
