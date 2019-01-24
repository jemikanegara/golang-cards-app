package main

func main() {
	// card := showCard()
	// fmt.Println(card)
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "New Cards")
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }

	// cards := newDeck()

	// PRINT FUNCTION
	// cards.print()

	// CHECK TO STRING VALUE
	// fmt.Println(cards.toString())

	// DEAL FUNCTION
	// hand, remainingCards := deal(cards, 10)
	// fmt.Println(hand)
	// fmt.Println(remainingCards)

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func newCard() string {
// 	return "Ace of Spades"
// }
