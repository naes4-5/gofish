package gofish

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var suits map[int]string = map[int]string{
	0: "♤",
	1: "♡",
	2: "♢",
	3: "♧",
}

func TakeTurn(players []Player, deck Deck) (logs string, err error) {
	if len(players) == 0 {
		return "", errors.New("No players to take the turn")
	}
	logs = ""
	for i, player := range players {
		choice := rand.IntN(len(players))
		for choice == i {
			choice = rand.IntN(len(players))
		}
		choiceRank := rand.IntN(13) + 1
		logs += fmt.Sprintf("Player %d has asked for a %d from player %d. ", i+1, choiceRank, choice+1)
		ind, amt, err := players[choice].handContains(choiceRank)
		if err != nil {
			return "", err
		}
		if ind == -1 {
			logs += fmt.Sprintf("Player %d does not have any %ds\n", choice+1, choiceRank)
			drawnCard, err := deck.drawCard()
			if err != nil {
				return "", err
			}
			players[i].Hand = append(player.Hand, drawnCard)
			continue
		}
		logs += fmt.Sprintf("Player %d has %d %ds and they were given given to player %d\n", choice+1, amt, choiceRank, i+1)
		removed := players[choice].Hand[ind : ind+amt]
		players[choice].Hand = append(players[choice].Hand[:ind], players[choice].Hand[ind+amt:]...)
		players[i].Hand = append(player.Hand, removed...)
		sortHand(&players[i])
		_, err = player.removeBooks()
		if err != nil {
			return "", err
		}
	}
	return logs, nil
}
