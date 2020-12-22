package day22

import (
	"fmt"
	"strconv"
	"strings"
)

func task2() {
	p1, p2 := getInputs()

	winner, p1, p2 := game(p1, p2, 1)

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

	fmt.Printf("\nDay 22 task 2: %s\n", message)
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

func game(p1Deck, p2Deck []int, level int) (int, []int, []int) {
	plays := make(map[string]struct{}, 0)
	p1Card, p2Card := 0, 0
	var err error
	round := 0
	for {
		round++
		fmt.Printf("------- Round %3d of Game %2d -------\nbeginning a new iteration with decks:\np1: %v\np2: %v\n\n", round, level, p1Deck, p2Deck)
		// if either of the decks is now empty, stop the game, we have won
		if len(p1Deck) == 0 {
			fmt.Printf("p1's deck is now empty, returning 2 as player 2 won, %v, %v\n", p1Deck, p2Deck)
			return 2, p1Deck, p2Deck
		}
		if len(p2Deck) == 0 {
			fmt.Printf("p2's deck is now empty, returning 1 as player 1 won, %v, %v\n", p1Deck, p2Deck)
			return 1, p1Deck, p2Deck
		}

		// check if we've played this before
		if _, ok := plays[playString(p1Deck, p2Deck)]; ok {
			// player 1 won
			fmt.Printf("we've played this before, returning 1, %v, %v\n", p1Deck, p2Deck)
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
			fmt.Printf("new recursive game!\np1card: %2d, len p1deck: %2d\np1deck: %v\n\np2card: %2d, len p2deck: %2d\np2deck: %v\n\n",
				p1Card,
				len(p1Deck),
				p1Deck,
				p2Card,
				len(p2Deck),
				p2Deck,
			)
			p1Copy := make([]int, p1Card)
			p2Copy := make([]int, p2Card)
			copy(p1Copy, p1Deck)
			copy(p2Copy, p2Deck)

			winner, _, _ := game(p1Copy, p2Copy, level+1)

			fmt.Printf("returned from Game %d\np1deck: %v\np2deck: %v\np1card: %d\np2card: %d\nwinner: %d\n", level+1, p1Deck, p2Deck, p1Card, p2Card, winner)
			switch winner {
			case 1:
				p1Deck = sliceAdd(p1Deck, p1Card, p2Card)
			case 2:
				p2Deck = sliceAdd(p2Deck, p2Card, p1Card)
			default:
				panic(fmt.Sprintf("unexpected winner value: %d", winner))
			}
			fmt.Printf("player %d won Round %d of game %d\np1: %v\np2: %v\n\n", winner, round, level, p1Deck, p2Deck)
			continue
		}

		// no winner yet for other reasons, so let's compare the drawn cards
		// if player 1 won, then put the two cards into player one's deck
		if p1Card > p2Card {
			fmt.Printf("no winner yet, p1 won with higher card\n")
			p1Deck = sliceAdd(p1Deck, p1Card, p2Card)
		} else {
			fmt.Printf("no winner yet, p2 won with higher card\n")
			p2Deck = sliceAdd(p2Deck, p2Card, p1Card)
		}
		fmt.Printf("cards:\np1: %2d\np2: %2d\n", p1Card, p2Card)

		fmt.Printf("new decks:\np1: %v\np2: %v\n", p1Deck, p2Deck)
	}
}
