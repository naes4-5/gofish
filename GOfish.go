package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
	"sort"
	//"log"
	"os"
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

func printHands(players []player_t) {
	for _, player := range players {
		for i, card:= range player.hand {
			fmt.Printf("%d of %s\n", card.rank, card.suit)
			if i == len(player.hand) - 1 {
				fmt.Printf("\n")
			}
		}
	}
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

func (deck *deck_t) startGame(handSize int, players ...*player_t) ([]player_t, error) {
	if len(players) * handSize > deck.cardsLeft {
		return []player_t{}, errors.New(fmt.Sprintf("Too many players for handsize"))
	} else if len(players) < 2 {
		return []player_t{}, errors.New(fmt.Sprintf("Not enough players to play"))
	}
	playerList := make([]player_t, len(players))
	for _, player := range players {
		for i := 0; i < handSize; i++ {
			card, err := deck.drawCard()
			if err != nil {
				return []player_t{}, err
			}
			player.hand = append(player.hand, card)
		}
		sortHand(player)
		playerList = append(playerList, *player)
	}
	return playerList, nil
}

func (player *player_t) bookCheck() []int {
	bookedRanks := make([]int, 1)
	c := 1
	for i := 0; i < len(player.hand)-1; i++ {
		if player.hand[i].rank != player.hand[i+1].rank {
			c = 1
			continue
		}
		c++
		if c == 4 {
			bookedRanks = append(bookedRanks, player.hand[i].rank)
		}
	}
	return bookedRanks
}

func (player *player_t) removeBooks(ranks []int) int {
	for _, rank := range ranks {
		for i := 0; i < len(player.hand); i++ {
			if player.hand[i].rank == rank {
				player.hand = append(player.hand[:i], player.hand[i+4:]...)
				player.books++
				break
			}
		}
	}
	return len(ranks)
}

func main() {
	deck := newDeck()
	p1 := player_t {hand: []card_t{}, books: 0}
	p2 := player_t {hand: []card_t{}, books: 0}
	p3 := player_t {hand: []card_t{}, books: 0}
	p4 := player_t {hand: []card_t{}, books: 0}
	
	handSize := 5
	players, err := deck.startGame(handSize, &p1, &p2, &p3, &p4)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	printHands(players)
}

