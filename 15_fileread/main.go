package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileReader()
}

func fileReader() {
	file, err := os.Open("small.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 20)

	for {
		n, err := file.Read(buf)
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
