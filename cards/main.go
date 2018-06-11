package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	//cards.print()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	fmt.Println("Printing out the rest of the deck")
	remainingDeck.print()

	fmt.Println(cards.toByteSlice())
}
