package main

import "fmt"

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
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
