package main

import (
	"fmt"
	"strings"
 "io/ioutil"
)

type deck []string
// Create a new type of "deck"
// which is a slice of strings


func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _,suit := range cardSuits {
		for _,value := range cardValues {
					cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func newCard() string {
	return "Five of Diamonds"
}

func (d deck) print() {
	for i,card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}


func (d deck) saveToFile(filename string) error {
	// 0666 is default permissions on the file - anyone can read and write
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
