package main

func main() {
	// var card string = "Ace of spades"
	// card := newCard()
	// fmt.Println(card)
	// printState()

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
