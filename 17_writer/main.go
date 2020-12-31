package main

import (
	"os"
)

func main() {
	f, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte("writing some string into the file using write method\n"))
}
