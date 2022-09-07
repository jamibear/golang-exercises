package main

// Project: Playing Cards
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
