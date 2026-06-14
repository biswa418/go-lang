package main

import "fmt"

// Creating a new type of 'deck'
type deck []string

// creates and returns a list of playing cards -- Array of strings
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

// log out the contents of a deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deals the cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
