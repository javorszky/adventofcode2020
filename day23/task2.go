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
}

func (b bigOof) step() bigOof {
	// get the index of current
	idxCurrent := b.whereIs[b.current]

	p1, p2, p3 := b.whatsOn[((idxCurrent)%b.length)+1], b.whatsOn[((idxCurrent+1)%b.length)+1], b.whatsOn[((idxCurrent+2)%b.length)+1]

	// let's find the next one
	next := b.current - 1
	for {
		// if it was supposed to be 0, then we loop around, and start from the top.
		if next == 0 {
			next = b.length
		}
		// if the next is none of the picked up values, break, we found it!
		if next != p1 && next != p2 && next != p3 {
			break
		}
		// otherwise subtract one more, and try again.
		next--
	}

	// where is next
	idxNext := b.whereIs[next]

	fmt.Printf("whatson: %#v\n"+
		"current: %d: %d\n"+
		"next: %d: %d\n"+
		"p1, p2, p3: %d %d %d\n",
		b.whatsOn,
		idxCurrent, b.current,
		idxNext, next,
		p1, p2, p3,
	)

	// if we're moving right, to a higher
	if idxCurrent-idxNext > 0 {
		fmt.Printf("we're shifting things right, as idx next is lower\n")
		// if we can partition the elements moved into two groups, do it.
		if idxCurrent-idxNext > 3 {
			fmt.Printf("the gap is big\nidxcurrent is %d\n\n", idxCurrent)
			// make sure we calculate the modulo index for the first three elements. This was we don't need to calculate it
			// it for hundreds of thousands of elements.
			for i := idxCurrent; i > idxCurrent-3; i-- {
				shift := ((i + 2) % b.length) + 1
				fmt.Printf("shifting (with mod) %d -> %d\n", i, shift)
				n := b.whatsOn[i]
				b.whatsOn[shift] = n
				b.whereIs[n] = shift
			}

			// all others will not wrap around, we are okay to use the straight indexes without modulo.
			for j := idxCurrent - 3; j > idxNext; j-- {
				shift := j + 3
				fmt.Printf("shifting (no mod) %d -> %d\n", j, shift)
				n := b.whatsOn[j]
				b.whatsOn[shift] = n
				b.whereIs[n] = shift
			}
		} else {

			fmt.Printf("the gap is smol\n")
			// if the gap is small, we might as well just calculate the modulos.
			for j := idxCurrent; j > idxNext; j-- {
				shift := ((j + 2) % b.length) + 1

				fmt.Printf("shifting (with mod) %d -> %d\n", j, shift)
				n := b.whatsOn[j]
				b.whatsOn[shift] = n
				b.whereIs[n] = shift
			}
		}

		idxp1 := ((idxNext) % b.length) + 1
		idxp2 := ((idxNext + 1) % b.length) + 1
		idxp3 := ((idxNext + 2) % b.length) + 1
		fmt.Printf("picked idxes: %d, %d, %d\n", idxp1, idxp2, idxp3)
		b.whatsOn[idxp1] = p1
		b.whereIs[p1] = idxp1

		b.whatsOn[idxp2] = p2
		b.whereIs[p2] = idxp2

		b.whatsOn[idxp3] = p3
		b.whereIs[p3] = idxp3
	} else {
		fmt.Printf("we're shifting things left, as idx next is higher\n")

		fmt.Printf("the gap is smol\n")
		// if the gap is small, we might as well just calculate the modulos.
		for j := idxCurrent + 4; j <= idxNext; j++ {
			shift := j - 3

			fmt.Printf("shifting (with mod) %d -> %d\n", j, shift)
			n := b.whatsOn[j]
			b.whatsOn[shift] = n
			b.whereIs[n] = shift
		}

		idxp1 := idxNext - 2
		idxp2 := idxNext - 1
		idxp3 := idxNext
		fmt.Printf("picked idxes: %d, %d, %d\n", idxp1, idxp2, idxp3)
		b.whatsOn[idxp1] = p1
		b.whereIs[p1] = idxp1

		b.whatsOn[idxp2] = p2
		b.whereIs[p2] = idxp2

		b.whatsOn[idxp3] = p3
		b.whereIs[p3] = idxp3
	}

	// after having moved numbers, where's the current again?
	idxCurrentMoved := b.whereIs[b.current]
	fmt.Printf("\n\ncurrent is %d at %d\nwhatson: %#v\nnext idx should be %d\n", b.current, idxCurrentMoved, b.whatsOn, ((idxCurrentMoved)%b.length)+1)

	b.current = b.whatsOn[((idxCurrentMoved)%b.length)+1]
	return b
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
		length:  cups,
		current: whatisthere[1],
		whatsOn: whatisthere,
		whereIs: whereisit,
	}
}
