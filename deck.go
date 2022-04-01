package main

import "fmt"

type deck []string

// anytime someone returns newDeck they are gonna return a type of value of deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print () {
	for i, card := range d {
		fmt.Println(i, card)
	}
}