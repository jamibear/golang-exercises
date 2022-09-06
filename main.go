package main

import "fmt"

// Project: Playing Cards
func main() {
	cards := newDeck()

	fmt.Println(cards.toString())
}
