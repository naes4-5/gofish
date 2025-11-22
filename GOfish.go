package gofish

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
)

var suits map[int]string = map[int]string{
	0: "♤",
	1: "♡",
	2: "♢",
	3: "♧",
}

type Card struct {
	suit string
	rank int
}

func takeTurn(players []Player, deck Deck) (logs string, err error) {
	if len(players) == 0 {
		return "", errors.New("No players to take the turn")
	}
	for i, player := range players {
		choice := rand.IntN(len(players))
		for choice == i {
			choice = rand.IntN(len(players))
		}
		choiceRank := rand.IntN(13) + 1
		ind, amt, err := players[choice].handContains(choiceRank)
		if err != nil {
			return "", err
		}
		if ind == -1 {
			drawnCard, err := deck.drawCard()
			if err != nil {
				return "", err
			}
			player.hand = append(player.hand, drawnCard)
			continue
		}
		//now remove cards that were found in the chosen player's hand and add them to the current player's hand
		player.removeBooks()
	}
	return "", nil
}

func main() {
	deck := newDeck()
	p1 := Player{hand: []Card{}, books: 0}
	p2 := Player{hand: []Card{}, books: 0}
	p3 := Player{hand: []Card{}, books: 0}

	handSize := 5
	players, err := deck.startGame(handSize, &p1, &p2, &p3)
	if err != nil {
		log.Fatal(err)
	}
	printHands(players)
	removed, err := players[len(players)-1].removeBooks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n%v\n", removed, players[len(players)-1].hand)
}
