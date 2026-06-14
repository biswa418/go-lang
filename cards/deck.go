package main

import (
	"fmt"
	"os"
	"strings"
)

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

// helper function -- deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to File
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

// read the file
func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
