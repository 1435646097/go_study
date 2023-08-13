package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/common_mistakes")
	if err != nil {
		fmt.Println(err)
	}
	err3 := db.Ping()
	if err3 != nil {
		fmt.Println(err3)
	}
	// p := new(Person)
	// p.id = 1542142896440901634
	// p.FindOne()
	// //更新
	// p.id = 1542155000917057536
	// p.name = "你好世界"
	// p.insert()
	// p, err2 := findById(1542140936379777026)
	// if err2 != nil {
	// fmt.Println(err2)
	// }
	// fmt.Println(p)
	fmt.Printf("findAll(): %v\n", findAll())
}
