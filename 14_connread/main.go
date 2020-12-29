package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	connReader()
}

func connReader() {
	conn, err := net.Dial("tcp", "google.com:80")

	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, "GET HTTP 1.0 \r\n\r\n")
	buf := make([]byte, 50)

	for {
		n, err := conn.Read(buf)
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
