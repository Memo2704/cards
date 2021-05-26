package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard() // Creating a variable without the type
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
