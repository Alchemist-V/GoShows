package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// MyWriter inteface
type MyWriter interface {
	Write(p []byte) (n int, err error)
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)

	// Any IO shall implement reader and closer interface.
	resp.Body.Read(bs)
	io.Copy(os.Stdout, VBody.Body)

}
