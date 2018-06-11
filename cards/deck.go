package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

//Create a new type of 'deck'
//which is a slice of strings

type deck []string

func check(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toByteSlice(s string) []byte {
	return []byte(s)
}

func (d deck) toByteSlice() []byte {
	return toByteSlice(d.toString())
}

func (d deck) saveToFile(filename string) bool {
	err := ioutil.WriteFile("/tmp/"+filename, d.toByteSlice(), 0755)
	return check(err)
}

func newDeckFromFile(filename string) deck {
	b, err := ioutil.ReadFile("/tmp/" + filename)
	if check(err) {
		log.Fatal(err)
		fmt.Println("Error: ", err)
		return newDeck()
	}
	//Convert from []byte to deck
	str := string(b)
	bs := splitStringOnComma(str)
	return deck(bs)
}

func splitStringOnComma(str string) []string {
	return strings.Split(str, ",")
}

func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		randomNum := r.Int() % (52 - i)
		fmt.Println("The random integer: ", randomNum)
		d[i], d[i+randomNum] = d[i+randomNum], d[i]
	}
}
