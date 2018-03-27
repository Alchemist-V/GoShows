package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
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

// writes to a file on disk
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Reads content from file and returns a deck object.
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	str := strings.Split(string(bs), ",")
	return deck(str)
}

// shuffles the cards in the deck
func (d deck) shuffle() {
	// getting a new seed everytime to ensure non repeatitive random number generation.
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
