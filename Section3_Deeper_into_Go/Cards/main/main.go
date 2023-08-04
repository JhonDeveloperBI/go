package main

func main() {
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}
