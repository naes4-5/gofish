package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
	"sort"
	//"log"
	//"os"
)

var suits map[int]string = map[int]string {
	0: "♤", 
	1: "♡", 
	2: "♢", 
	3: "♧",
}

type card_t struct {
	suit string;
	rank int;
}

type deck_t struct {
	cards        [4][]card_t;
	cardsLeft    int;
	cardsPerSuit int;
}

type player_t struct {
	hand  []card_t;
	books int;
}

func sortHand(player *player_t) {
	sort.Slice(player.hand, func(i, j int) bool {
		return player.hand[i].rank < player.hand[j].rank
	})
}

func newDeck() deck_t {
	var deck deck_t
	deck.cardsPerSuit = 13
	deck.cards = [4][]card_t{}
	for s := 0; s < len(suits); s++ {
		deck.cards[s] = make([]card_t, deck.cardsPerSuit)
		for i := 0; i < deck.cardsPerSuit; i++ {
			deck.cards[s][i] = card_t {suit: suits[s], rank: i+1}
		}
	}
	deck.cardsLeft = 52
	return deck
}

func (deck *deck_t) drawCard() (card_t, error) {
	if deck.cardsLeft <= 0 {
		return card_t {}, errors.New(fmt.Sprintf("no more cards to draw"))
	}

	rsuit := rand.IntN(4)
	for len(deck.cards[rsuit]) == 0 {
		rsuit = rand.IntN(4)
	}
	i := rand.IntN(len(deck.cards[rsuit]))
	ret := deck.cards[rsuit][i]
	
	deck.cards[rsuit] = append(deck.cards[rsuit][:i], deck.cards[rsuit][i+1:]...)
	deck.cardsLeft--
	return ret, nil
}

func (deck *deck_t) startGame(handSize int, players ...*player_t) error {
	if len(players) * handSize > deck.cardsLeft {
		return errors.New(fmt.Sprintf("Too many players for handsize"))
	} else if len(players) < 2 {
		return errors.New(fmt.Sprintf("Not enough players to play"))
	}
	for _, player := range players {
		for i := 0; i < handSize; i++ {
			card, err := deck.drawCard()
			if err != nil {
				return err
			}
			player.hand = append(player.hand, card)
		}
		sortHand(player)
	}
	return nil
}

func (player *player_t) bookCheck() int

func main() {
	deck := newDeck()
	p1 := player_t {hand: []card_t{}, books: 0}
	p2 := player_t {hand: []card_t{}, books: 0}
	p3 := player_t {hand: []card_t{}, books: 0}
	
	deck.startGame(5, &p1, &p2, &p3)
	for i := 0; i < 5; i++ {
		card := p1.hand[i]
		fmt.Printf("%d of %s\n", card.rank, card.suit)
	}
	fmt.Printf("\n\n")
	for i := 0; i < 5; i++ {
		card := p2.hand[i]
		fmt.Printf("%d of %s\n", card.rank, card.suit)
	}
	fmt.Printf("\n\n")
	for i := 0; i < 5; i++ {
		card := p3.hand[i]
		fmt.Printf("%d of %s\n", card.rank, card.suit)
	}
}

