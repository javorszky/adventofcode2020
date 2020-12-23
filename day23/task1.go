package day23

import (
	"fmt"
	"strconv"
	"strings"
)

func task1() {
	start := getInputs()
	for i := 0; i < 100; i++ {
		start = round(start)
	}

	fmt.Printf("\nDay 23 task 1: after 100 rounds of Crab Cups the final order is going to be '%s'\n", orderOfCups(start))
}

// rotateCircleBy does one thing: takes the first element off of slice, sticks it in the end, and returns it.
func rotateCircleBy(in []int, n int) []int {
	if n < 1 {
		return in
	}
	realRotation := n % len(in)
	return append(in[realRotation:], in[0:realRotation]...)
}

// removeThreeCups will remove three elements from the slice. The elements removed are always the 1,2,3 indexed parts,
// because in our solution the 0 idx element is the "current" cup.
func removeThreeCups(in []int) ([]int, []int) {
	snip := make([]int, 3)
	short := make([]int, len(in)-3)
	copy(snip, in[1:4])
	copy(short, append(in[0:1], in[4:]...))
	return short, snip
}

// findIndexOfNextCurrentCup will return the index of the next current cup.
func findIndexOfNextCurrentCup(in []int) int {
	// current cup is the 0 idx one
	current := in[0]
	existenceHelper := make(map[int]int, 0)
	highest := 0
	for idx, v := range in {
		if v > highest {
			highest = v
		}
		existenceHelper[v] = idx
	}

	trying := current - 1
	for {
		if trying < 0 {
			trying = highest
		}
		idx, ok := existenceHelper[trying]
		if ok {
			return idx
		}
		trying--
	}
}

// insertSnippetIntoAt will insert a slice into another slice at a certain index. This is somewhat counterintuitive,
// because given a slice of [1 2 3 4], index 3 is the number 4. Inserting [5 6] at index 3 will result in [5 6]
// becoming the thing at index 3, and whatever was at index 3 (the number 4), comes at the end, so the end result is
// [1 2 3 5 6 4].
//
// To insert something at the front, use idx 0.
// To insert something at the end, use idx len(short).
func insertSnippetIntoAt(short, snippet []int, idx int) []int {
	if idx > len(short) {
		panic(fmt.Sprintf("can't insert stuff at index %d into a slice with lenght of %d, out of bounds", idx, len(short)))
	}

	head := make([]int, len(short[0:idx]))
	tail := make([]int, len(short[idx:]))
	copy(tail, short[idx:])
	copy(head, short[0:idx])

	intermediary := append(head, snippet...)
	intermediary = append(intermediary, tail...)

	return intermediary
}

// round will play a round of Crab Cups per the rules of Advent Of Code 2020 Day 23 Task 1.
func round(in []int) []int {

	// n=0 idx is the "current" cup
	// step 1: remove the snippet
	short, snippet := removeThreeCups(in)
	// n+1,2,3 are picked up and removed from the circle
	nextCupIDX := findIndexOfNextCurrentCup(short)

	newState := insertSnippetIntoAt(short, snippet, nextCupIDX+1)

	// rotate short so that next cup is at idx 0
	newState = rotateCircleBy(newState, 1)

	// add back the snippet

	return newState
}

func orderOfCups(in []int) string {
	idx := 0
	for i, v := range in {
		if v == 1 {
			idx = i
			break
		}
	}
	ordered := rotateCircleBy(in, idx)[1:]
	var sb strings.Builder
	for _, v := range ordered {
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}
