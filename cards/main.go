package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, remainingHand := deal(cards, 5)
	fmt.Println(hand.toString())
	fmt.Println(remainingHand.toString())
}

// returns new card
func newCard() string {
	return "5 of Diamonds"
}
