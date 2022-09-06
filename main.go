package main

// Project: Playing Cards
func main() {
	cards := newDeck()

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
