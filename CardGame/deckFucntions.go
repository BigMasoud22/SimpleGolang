package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	types := []string{"Clubs", "Hearts", "Spades", "Diamonds"}
	cardslevels := []string{"Ace of ", "Two of ", "Three of ", "Four of ", "Five of ", "Six of ", "Seven of ", "Eight of ", "Nine of ", "Ten of ", "Jack of ", "Queen of ", "King of "}
	returnDeck := deck{}
	for _, Type := range types {
		for _, Card := range cardslevels {
			returnDeck = append(returnDeck, Card+Type)
		}
	}
	return returnDeck
}

func (dealed deck) shuffle() deck {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	Deck := dealed
	for index := range Deck {
		randomIndex := random.Intn(len(Deck) - 1)
		Deck[index], Deck[randomIndex] = Deck[randomIndex], Deck[index]
	}
	return Deck
}

func deal(cards deck, count int) (deck, deck) {
	return cards[:count], cards[count:]
}

func (Data deck) writeOnFile(filename string) (string, error) {

	myStrings := []string(Data)
	isolated := strings.Join(myStrings, ",")
	return filename, ioutil.WriteFile(filename, []byte(isolated), 0666)
}

func readFromFile(filename string) (deck, error) {

	data, err := ioutil.ReadFile(filename)
	myDeck := deck(strings.Split(string(data), ","))
	return myDeck, err
}
