package main

import (
	"fmt"
)

func main() {
	var str String
	str.Write([]byte("Hello world!"))
	fmt.Println(str)
}
