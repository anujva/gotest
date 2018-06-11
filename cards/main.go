package main

func main() {
	cards := newDeck()
	//cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	//fmt.Println("Printing out the rest of the deck")
	// remainingDeck.print()
	// cards.saveToFile("cardsOnDisk.txt")
	//fmt.Println(cards.toByteSlice())
	//fmt.Println("")
	//fmt.Println("")
	//fmt.Println("")
	//fmt.Println("")
	//fmt.Println("")
	//fmt.Println("Printing out the deck read from the file: ")
	//newDeckFromFile("cardsOnDisk.txt").print()
	cards.shuffle()
	cards.print()
}
