package main

import (
	"fmt"
	"io"
	"os"
)

type String string

func (s *String) Write(p []byte) (n int, err error) {
	*s += String(string(p))
	return len(p), nil
}

func WriteInString() {
	var str String
	str.Write([]byte("Hello"))
	fmt.Println(str)
}

func SayHello(w io.Writer) {
	w.Write([]byte("Hello"))
}

func EchoSayHello() {
	var str String
	SayHello(&str)
}

func EchoSayHelloInStdout() {
	SayHello(os.Stdout)
}

func EchoSayHelloInFile() {
	file, _ := os.Create("hello.txt")
	defer file.Close()
	SayHello(file)
}
