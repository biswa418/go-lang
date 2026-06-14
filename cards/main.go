package main

// import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := newCard()
	// fmt.Println(card)
	// printState()

	cards := newDeckFromFile("my_cards")

	cards.print()
}
