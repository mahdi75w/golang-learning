package main

import (
	"fmt"
	"io"
	"os"
)

type AdvanceString struct {
	data string
	pos  int
}

func (s *AdvanceString) Read(b []byte) (int, error) {
	n := copy(b, s.data[s.pos:])
	s.pos += n

	var err error
	if s.pos >= len(s.data) {
		err = io.EOF
	}
	return n, err
}

func (s *AdvanceString) Write(b []byte) (int, error) {
	s.data += string(b)
	return len(b), nil
}

func SampleRead() {
	var astr AdvanceString
	astr.Write([]byte("Hello World"))
	buffer := make([]byte, 20)
	n, err := astr.Read(buffer)
	fmt.Println("Bytes read:", n)
	fmt.Println("error:", err)
	fmt.Println("buffer:", string(buffer))
}

func ReadWithLowBuffer() {
	var astr AdvanceString
	astr.Write([]byte("Hello World"))
	buffer := make([]byte, 3)
	n, err := astr.Read(buffer)
	fmt.Println("Bytes read:", n)
	fmt.Println("error:", err)
	fmt.Println("buffer:", string(buffer))
}

func ReadFullOfStream() {
	var astr AdvanceString
	astr.Write([]byte("Hello World"))
	buffer := make([]byte, 3)
	var n int
	var err error
	for err != io.EOF { // WARNING : buffer reuse in every circle, so maybe has old data in it per circle21
		n, err = astr.Read(buffer)
		fmt.Print(string(buffer[:n]), "/")
	}
}

func ReadFromInput() {
	buffer := make([]byte, 10)
	n, _ := os.Stdin.Read(buffer)
	fmt.Println("Buffer:", string(buffer[:n]))
}
