package main

import (
	"fmt"
	"errors"
	"math/rand/v2"
	"sort"
	"log"
)

var suits map[int]string = map[int]string {
	0: "♤", 
	1: "♡", 
	2: "♢", 
	3: "♧",
}

type Card struct {
	suit string;
	rank int;
}

type Deck struct {
	cards        [4][]Card;
	cardsLeft    int;
	cardsPerSuit int;
}

type Player struct {
	hand  []Card;
	books int;
}

func sortHand(player *Player) {
	sort.Slice(player.hand, func(i, j int) bool {
		return player.hand[i].rank < player.hand[j].rank
	})
}

func (player *Player) handContains(rank int) (int, error) {
	if len(player.hand) == 0 {
		return -1, errors.New("No cards in hand")
	}
	for i, card := range player.hand {
		if card.rank == rank {
			return i, nil
		}
	}
	return -1, nil
}

func printHands(players []Player) {
	for _, player := range players {
		for i, card:= range player.hand {
			fmt.Printf("%d of %s\n", card.rank, card.suit)
			if i == len(player.hand) - 1 {
				fmt.Printf("\n")
			}
		}
	}
}

func newDeck() Deck {
	var deck Deck
	deck.cardsPerSuit = 13
	deck.cards = [4][]Card{}
	for s := 0; s < len(suits); s++ {
		deck.cards[s] = make([]Card, deck.cardsPerSuit)
		for i := 0; i < deck.cardsPerSuit; i++ {
			deck.cards[s][i] = Card {suit: suits[s], rank: i+1}
		}
	}
	deck.cardsLeft = 52
	return deck
}

func (deck *Deck) drawCard() (Card, error) {
	if deck.cardsLeft <= 0 {
		return Card {}, errors.New(fmt.Sprintf("no more cards to draw"))
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

func (deck *Deck) startGame(handSize int, players ...*Player) ([]Player, error) {
	if len(players) * handSize > deck.cardsLeft {
		return []Player{}, errors.New("Too many players for handsize")
	} else if len(players) < 2 {
		return []Player{}, errors.New("Not enough players to play")
	}
	playerList := make([]Player, len(players))
	for _, player := range players {
		startingCardsInHand := len(player.hand)
		for i := 0; i < handSize - startingCardsInHand; i++ {
			card, err := deck.drawCard()
			if err != nil {
				return []Player{}, err
			}
			player.hand = append(player.hand, card)
		}
		sortHand(player)
		playerList = append(playerList, *player)
	}
	return playerList, nil
}

func (player *Player) bookCheck() []int {
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

func (player *Player) removeBooks(ranks []int) int {
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
	p1 := Player {hand: []Card{}, books: 0}
	p2 := Player {hand: []Card{}, books: 0}
	p3 := Player {hand: []Card{}, books: 0}
	p4 := Player {hand: []Card{}, books: 0}
	p5 := Player {
		hand: []Card {
			Card {
				rank: 5,
				suit: suits[3],
			}, 
			Card {
				rank: 8,
				suit: suits[1],
			},
		},
		books:0,
	}
	
	handSize := 5
	players, err := deck.startGame(handSize, &p1, &p2, &p3, &p4, &p5)
	if err != nil {
		log.Fatal(err)
	}
	printHands(players)
}

