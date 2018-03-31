package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Holla"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// not allowed, interface cannot be receiver.
// func (b bot) print() {
// 	fmt.Println(b.getGreeting())
// }

// fautly impl without using interfaces
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
