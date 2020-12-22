package day22

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day22/input.txt"

// getInputs reads the input.txt file and returns the two decks. Player 1 is the first return, Player 2 is the second.
func getInputs() ([]int, []int) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// remove trailing newline and split along the double linebreak
	decks := strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
	playerOne := parseDeck(decks[0])
	playerTwo := parseDeck(decks[1])
	return playerOne, playerTwo
}

func parseDeck(s string) []int {
	deckLines := strings.Split(s, "\n")
	player := make([]int, 0)

	for idx, line := range deckLines {
		if idx == 0 {
			continue
		}

		i, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("cant turn string '%s' into int: %s", line, err))
		}
		player = append(player, i)
	}
	return player
}
