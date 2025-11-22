package main

import (
	"fmt"
	"github.com/naes4-5/gofish/game"
	"log"
)

func main() {
	deck := gofish.NewDeck()
	p1 := gofish.Player{Hand: []gofish.Card{}, Books: 0}
	p2 := gofish.Player{Hand: []gofish.Card{}, Books: 0}
	p3 := gofish.Player{Hand: []gofish.Card{}, Books: 0}

	handSize := 5
	players, err := deck.StartGame(handSize, &p1, &p2, &p3)
	if err != nil {
		log.Fatal(err)
	}
	gofish.PrintHands(players)
	gofish.TakeTurn(players, deck)
	fmt.Printf("Turn taken ->\n")
	gofish.PrintHands(players)
}
