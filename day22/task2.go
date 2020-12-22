package day22

import (
	"fmt"
	"strconv"
	"strings"
)

func task2() {
	p1, p2 := getInputs()

	winner, p1, p2 := game(p1, p2)

	message := ""
	switch winner {
	case 0:
		if len(p2) == 0 {
			message = fmt.Sprintf("player 1 won with a score of %d", calculateScore(p1))
		} else {
			message = fmt.Sprintf("player 2 won with a score of %d", calculateScore(p2))
		}
	case 1:
		message = fmt.Sprintf("player 1 won with a score of %d", calculateScore(p1))
	case 2:
		message = fmt.Sprintf("player 2 won with a score of %d", calculateScore(p2))
	default:
		panic(fmt.Sprintf("unexpected winner value: %d", winner))
	}

	fmt.Printf("Day 22 task 2: %s\n", message)
}

func playString(p1, p2 []int) string {
	var sb strings.Builder
	stringHelper := make([]string, 0)
	for _, v := range p1 {
		stringHelper = append(stringHelper, strconv.Itoa(v))
	}
	sb.WriteString(strings.Join(stringHelper, ","))
	sb.WriteString("|")
	stringHelper = make([]string, 0)
	for _, v := range p2 {
		stringHelper = append(stringHelper, strconv.Itoa(v))
	}
	sb.WriteString(strings.Join(stringHelper, ","))

	return sb.String()
}

func game(p1Deck, p2Deck []int) (int, []int, []int) {
	plays := make(map[string]struct{}, 0)
	p1Card, p2Card := 0, 0
	var err error
	for {
		// if either of the decks is now empty, stop the game, we have won
		if len(p1Deck) == 0 {
			return 2, p1Deck, p2Deck
		}
		if len(p2Deck) == 0 {
			return 1, p1Deck, p2Deck
		}

		// check if we've played this before
		if _, ok := plays[playString(p1Deck, p2Deck)]; ok {
			// player 1 won
			return 1, p1Deck, p2Deck
		}

		// save the play as a unique play in the local map.
		plays[playString(p1Deck, p2Deck)] = struct{}{}

		// draw cards from both
		p1Card, p1Deck, err = slicePop(p1Deck)
		if err != nil {
			panic(fmt.Sprintf("could not pop card from p1's deck %#v: %s", p1Deck, err))
		}
		p2Card, p2Deck, err = slicePop(p2Deck)
		if err != nil {
			panic(fmt.Sprintf("could not pop card from p2's deck %#v: %s", p2Deck, err))
		}

		// If both players have at least as many cards remaining in their deck as the value of the card they just drew,
		// the winner of the round is determined by playing a new game of Recursive Combat (see below).
		if p1Card <= len(p1Deck) && p2Card <= len(p2Deck) {
			p1Copy := make([]int, p1Card)
			p2Copy := make([]int, p2Card)
			copy(p1Copy, p1Deck)
			copy(p2Copy, p2Deck)

			winner, _, _ := game(p1Copy, p2Copy)

			switch winner {
			case 1:
				p1Deck = sliceAdd(p1Deck, p1Card, p2Card)
			case 2:
				p2Deck = sliceAdd(p2Deck, p2Card, p1Card)
			default:
				panic(fmt.Sprintf("unexpected winner value: %d", winner))
			}
			continue
		}

		// no winner yet for other reasons, so let's compare the drawn cards
		// if player 1 won, then put the two cards into player one's deck
		if p1Card > p2Card {
			p1Deck = sliceAdd(p1Deck, p1Card, p2Card)
		} else {
			p2Deck = sliceAdd(p2Deck, p2Card, p1Card)
		}
	}
}
