package gofish

import (
	"errors"
	"fmt"
	"sort"
)

type Player struct {
	Hand  []Card
	Books int
}

func sortHand(player *Player) {
	sort.Slice(player.Hand, func(i, j int) bool {
		return player.Hand[i].Rank < player.Hand[j].Rank
	})
}

func (player *Player) handContains(rank int) (firstIndex int, numCardsOfRank int, err error) {
	if len(player.Hand) == 0 {
		return -1, 0, errors.New("No cards in hand")
	}
	index := -1
	amt := 0
	for i, card := range player.Hand {
		if card.Rank == rank {
			index = i
			for _, crad := range player.Hand[i:] {
				if !(crad.Rank == rank) {
					break
				}
				amt++
			}
			return index, amt, nil
		}
	}
	return index, amt, nil
}

func PrintHands(players []Player) {
	for _, player := range players {
		for i, card := range player.Hand {
			fmt.Printf("%d of %s\n", card.Rank, card.Suit)
			if i == len(player.Hand)-1 {
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
			if len(player.Hand) > 4 {
				combined := append(player.Hand[:ind], player.Hand[ind+4:]...)
				player.Hand = []Card{}
				player.Hand = append(player.Hand, combined...)
			} else {
				player.Hand = []Card{}
			}
			removed = append(removed, i)
		}
	}
	player.Books += len(removed)
	return removed, nil
}
