package main

import (
	"time"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
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

func newDeckFromFile(filename string) deck {
	// if no err, err will be nil (=null)
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
			fmt.Println("Error:", err)
			// any argument beside 0 indicates sth went wrong
			os.Exit(1)
	}
	s := strings.Split(string(bs), ",") //slice of strings
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
