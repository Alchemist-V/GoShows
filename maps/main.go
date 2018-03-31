package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "0000ff",
		"white": "000000",
	}

	// fmt.Println(colors)
	print(colors)

	// other declaration options:

	// colors2 := make(map[string]string)
	// colors2["blue"] = "00ff00"

	// fmt.Println(colors2)

	// delete(colors, "red")
	// fmt.Println(colors)

	// iterate over map
}

func print(m map[string]string) {
	for color, hex := range m {
		fmt.Println(color + " -> " + hex)
	}
}
