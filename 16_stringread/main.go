package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	stringsReader()
}

func stringsReader() {
	s := strings.NewReader("hello this my first attempt at creating strings reader")
	buf := make([]byte, 5)

	for {
		n, err := s.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
