package main

import "fmt"

func main() {
	//Let's play the game :D
	mainDeck := newDeck().shuffle()
	players := []Player{Player{Name: "player 1"}, Player{Name: "player 2"}, Player{Name: "player 3"}, Player{Name: "player 4"}}
	for j := range players {
		cards, remainingCards := deal(mainDeck, 13)
		players[j].Cards = append(players[j].Cards, cards...)
		mainDeck = remainingCards // Update the mainDeck with the remaining cards
	}
	//Showing players hands
	for _, player := range players {
		fmt.Println(player.Hand())
		fmt.Println()
	}
}
