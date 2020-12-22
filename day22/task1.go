package day22

import (
	"errors"
	"fmt"
)

func task1() {
	_, _ = getInputs()
}

// slicePop takes in a slice, pops the first value off, and returns the value, and the slice without the first value.
// Returns error in case slice is nil or less than one element long.
func slicePop(in []int) (int, []int, error) {
	if in == nil {
		return 0, nil, errors.New("slicePop: passed in nil slice")
	}

	if len(in) < 1 {
		return 0, nil, errors.New("slicePop: passed in empty slice")
	}
	return in[0], in[1:], nil
}

// sliceAdd will take a slice, and a number of ints, and returns a slice where the additional elements are added to the
// end of the original slice.
func sliceAdd(in []int, elements ...int) []int {
	for _, el := range elements {
		in = append(in, el)
	}
	return in
}

// play will take two slices of ints, pop the first one off, compare them, and add both of them to the winning slice to
// the bottom.
func play(playerOne, playerTwo []int) ([]int, []int) {
	p1Card, p1Deck, err := slicePop(playerOne)
	if err != nil {
		panic(fmt.Sprintf("could not pop card from p1's deck %#v: %s", playerOne, err))
	}
	p2Card, p2Deck, err := slicePop(playerTwo)
	if err != nil {
		panic(fmt.Sprintf("could not pop card from p2's deck %#v: %s", playerTwo, err))
	}
	// if player 1 won, then put the two cards into player one's deck
	if p1Card > p2Card {
		p1Deck = sliceAdd(p1Deck, p1Card, p2Card)
	} else {
		p2Deck = sliceAdd(p2Deck, p2Card, p1Card)
	}

	return p1Deck, p2Deck
}

func calculateScore(in []int) int {
	i := len(in)
	accumulator := 0
	for _, v := range in {
		accumulator = accumulator + v*i
		i--
	}
	return accumulator
}
