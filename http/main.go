package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)

	// Any IO shall implement reader and closer interface.
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
