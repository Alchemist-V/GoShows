package main

import (
	"fmt"
	"io"
	"os"
)

// MyReader interface
type MyReader interface {
	Read(b []byte) (n int, err error)
}

type myreader struct {
	fileName string
}

func main() {
	args := os.Args
	validateArgs(args)
	mr := myreader{
		fileName: args[1],
	}

	bs := make([]byte, 32*1024)
	mr.Read(bs)
}

func validateArgs(s []string) {
	if len(s) <= 0 {
		fmt.Println("Error: No args supplied to program")
		os.Exit(1)
	}
}

func (m myreader) Read(b []byte) (n int, err error) {
	fileName := m.fileName
	file, err := os.Open(fileName)
	io.Copy(os.Stdout, file)
	return len(b), nil
}
