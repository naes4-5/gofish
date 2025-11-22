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
	for range 5 {
		logs, err := gofish.TakeTurn(players, deck)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Turn taken -> logs: \n%s\n", logs)
		gofish.PrintHands(players)
		for i, player := range players {
			fmt.Printf("Player %d has %d books\n", i+1, player.Books)
		}
		fmt.Printf("End of turn \n\n")
	}
}
