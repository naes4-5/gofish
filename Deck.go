package gofish

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type Deck struct {
	cards        [4][]Card
	cardsLeft    int
	cardsPerSuit int
}

func newDeck() Deck {
	var deck Deck
	deck.cardsPerSuit = 13
	deck.cards = [4][]Card{}
	for s := 0; s < len(suits); s++ {
		deck.cards[s] = make([]Card, deck.cardsPerSuit)
		for i := 0; i < deck.cardsPerSuit; i++ {
			deck.cards[s][i] = Card{suit: suits[s], rank: i + 1}
		}
	}
	deck.cardsLeft = 52
	return deck
}

func (deck *Deck) drawCard() (Card, error) {
	if deck.cardsLeft <= 0 {
		return Card{}, errors.New(fmt.Sprintf("no more cards to draw"))
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
	if len(players)*handSize > deck.cardsLeft {
		return []Player{}, errors.New("Too many players for handsize")
	} else if len(players) < 2 {
		return []Player{}, errors.New("Not enough players to play")
	}
	playerList := []Player{}
	for _, player := range players {
		startingCardsInHand := len(player.hand)
		for i := 0; i < handSize-startingCardsInHand; i++ {
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
