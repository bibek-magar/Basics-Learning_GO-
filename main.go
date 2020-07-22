package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := setCardName("Ace of Spades")
	// fmt.Println(card)
	// cards := deck{setCardName("Five of Heart"), setCardName("Ace of spades")}
	// cards = append(cards, "Six of Diamond")

	// cards.print()
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("My Cards")

	// cards := newDeckFromFile("MyCards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

// func setCardName(name string) string {
// 	return name
// }
