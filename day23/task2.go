package day23

import (
	"fmt"
	"sort"
)

const (
	cupsInCircle = 1000000
	//moveCupsThisManyTimes = 10000000
	moveCupsThisManyTimes = 8
)

type bigOof struct {
	whatsOn, whereIs map[int]int
	current, length  int
	pickedUp         []int
}

func task2() {
	//start := getInputs()
	start := []int{5, 4, 3, 2, 1}
	bigOof := generateTenPlusNCups(start, 10-len(start))
	fmt.Printf("bigoof to start with is %v\n", bigOof)
	for i := 0; i < moveCupsThisManyTimes; i++ {
		bigOof = round(bigOof)
		fmt.Printf("after round %2d, big oof is %v\n", i, bigOof)
	}

	fmt.Printf("Day 23 Task 2: after teh crab played ten million rounds of crab cups with one million cups, the product of the two cups clockwise to cup \"1\" is %d\n", productOfCupsClockwiseToOne(bigOof))
}

func generateTenPlusNCups(in []int, n int) []int {
	maxHelper := make([]int, len(in))
	copy(maxHelper, in)
	sort.Ints(maxHelper)
	max := maxHelper[len(maxHelper)-1]
	for i := max + 1; i <= max+n; i++ {
		in = append(in, i)
	}
	return in
}

func productOfCupsClockwiseToOne(in []int) int {
	idx := 0
	l := len(in)
	for i, v := range in {
		if v == 1 {
			idx = i
			break
		}
	}
	return in[(idx+1)%l] * in[(idx+2)%l]
}

// this will also have the
func newBigOof(input []int, cups int) bigOof {
	whereisit := make(map[int]int, 0)
	whatisthere := make(map[int]int, 0)
	for idx, i := range input {
		whatisthere[idx+1] = i
		whereisit[i] = idx + 1
	}
	for i := len(whatisthere) + 1; i <= cups; i++ {
		whatisthere[i] = i
		whereisit[i] = i
	}

	return bigOof{
		length:   cups,
		current:  whatisthere[1],
		whatsOn:  whatisthere,
		whereIs:  whereisit,
		pickedUp: nil,
	}
}
