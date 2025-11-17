package main

type String string

func (s *String) Write(b []byte) (n int, err error) {
	*s += String(string(b))
	return len(b), nil
}
