package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fmt.Println(readString)
}
