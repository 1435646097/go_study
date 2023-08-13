package controller

import (
	"fmt"
	"net/http"
)

func registeHomeController() {
	http.HandleFunc("/hello", helloWord)
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
