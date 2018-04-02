package main

import (
	"fmt"
)

func main() {
	ints := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	for i, val := range ints {
		fmt.Println(val + " is " + numtype(i))
	}
}

// return string indicating if the number is even or odd.
func numtype(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}
