package main

func main() {
	cards := newDeck()

	// deal returns 2 values
	// first value will be assigned to hand
	// second value will be assigned to remainingCards
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
