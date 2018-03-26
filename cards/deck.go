package main

import (
	"fmt"
	"strings"
)

// Create a new card deck, which is slice of string.

type deck []string

// Prints out a given deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returns a new fresh deck of playing cards.
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
		"Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Deals a given handsize
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) dealt(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
