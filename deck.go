package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// new type of 'deck' which is a slice of strings
type deck []string

// generate a new deck
func newDeck() deck {
	cards := deck{}

	// create pair of slices with card combinations
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// create new deck from combining suits and values slices
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal cards of hand size
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Join a slice of strings
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// Save deck to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1 - log error and return and return a call to newDeck()
		// option 2 - log error quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// convert string to deck type
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
