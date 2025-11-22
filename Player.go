package gofish

import (
	"errors"
	"fmt"
	"sort"
)

type Player struct {
	hand  []Card
	books int
}

func sortHand(player *Player) {
	sort.Slice(player.hand, func(i, j int) bool {
		return player.hand[i].rank < player.hand[j].rank
	})
}

func (player *Player) handContains(rank int) (firstIndex int, numCardsOfRank int, err error) {
	if len(player.hand) == 0 {
		return -1, 0, errors.New("No cards in hand")
	}
	index := -1
	amt := 0
	for i, card := range player.hand {
		if card.rank == rank {
			index = i
			for _, crad := range player.hand[i:] {
				if !(crad.rank == rank) {
					break
				}
				amt++
			}
			return index, amt, nil
		}
	}
	return index, amt, nil
}

func printHands(players []Player) {
	for _, player := range players {
		for i, card := range player.hand {
			fmt.Printf("%d of %s\n", card.rank, card.suit)
			if i == len(player.hand)-1 {
				fmt.Printf("\n")
			}
		}
	}
}

func (player *Player) removeBooks() (booksRemoved []int, err error) {
	var removed []int
	for i := 1; i < 14; i++ {
		ind, amt, err := player.handContains(i)
		if err != nil {
			return nil, fmt.Errorf("Error in handContains(): %s", err)
		}
		if amt == 4 {
			if len(player.hand) > 4 {
				combined := append(player.hand[:ind], player.hand[ind+4:]...)
				player.hand = []Card{}
				player.hand = append(player.hand, combined...)
			} else {
				player.hand = []Card{}
			}
			removed = append(removed, i)
		}
	}
	player.books += len(removed)
	return removed, nil
}
