package main

import (
	"net/http"
	"study/src/web_server/controller"
)

func main() {
	controller.RegisteRoute()
	http.ListenAndServe(":8080", nil)
}
