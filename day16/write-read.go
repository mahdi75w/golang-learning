package main

import (
	"fmt"
	"io"
)

func ReadWriteSomething(rw io.ReadWriter) string {
	rw.Write([]byte("Go language is fun!"))
	buff := make([]byte, 50)
	n, _ := rw.Read(buff)
	return string(buff[:n])
}

func read_write() {
	var str AdvanceString
	result := ReadWriteSomething(&str)
	fmt.Println(result)
}
