package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expecting length of 48 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expecting first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expecting the last card of King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readNewDeck("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected 48 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
