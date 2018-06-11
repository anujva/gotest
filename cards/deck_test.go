package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(deck))
	}

	if !(deck[0] == "Ace of Spades" && deck[len(deck)-1] == "King of Hearts") {
		t.Error("Expected the first card to be Ace of Spades and last one to be King of Hearts")
		t.Errorf("Got %v, %v", deck[0], deck[len(deck)-1])
	}
}
