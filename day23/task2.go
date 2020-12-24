package day23

import (
	"fmt"
)

const (
	cupsInCircle          = 1000000
	moveCupsThisManyTimes = 10000000
	//moveCupsThisManyTimes = 8
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

	// if we're moving right, to a higher
	if idxCurrent-idxNext > 0 {
		// if we can partition the elements moved into two groups, do it.
		if idxCurrent-idxNext > 3 {
			// make sure we calculate the modulo index for the first three elements. This was we don't need to calculate it
			// it for hundreds of thousands of elements.
			for i := idxCurrent; i > idxCurrent-3; i-- {
				shift := ((i + 2) % b.length) + 1
				n := b.whatsOn[i]
				b.whatsOn[shift] = n
				b.whereIs[n] = shift
			}

			// all others will not wrap around, we are okay to use the straight indexes without modulo.
			for j := idxCurrent - 3; j > idxNext; j-- {
				shift := j + 3
				n := b.whatsOn[j]
				b.whatsOn[shift] = n
				b.whereIs[n] = shift
			}
		} else {
			// if the gap is small, we might as well just calculate the modulos.
			for j := idxCurrent; j > idxNext; j-- {
				shift := ((j + 2) % b.length) + 1
				n := b.whatsOn[j]
				b.whatsOn[shift] = n
				b.whereIs[n] = shift
			}
		}

		idxp1 := ((idxNext) % b.length) + 1
		idxp2 := ((idxNext + 1) % b.length) + 1
		idxp3 := ((idxNext + 2) % b.length) + 1
		b.whatsOn[idxp1] = p1
		b.whereIs[p1] = idxp1

		b.whatsOn[idxp2] = p2
		b.whereIs[p2] = idxp2

		b.whatsOn[idxp3] = p3
		b.whereIs[p3] = idxp3
	} else {
		// if the gap is small, we might as well just calculate the modulos.
		for j := idxCurrent + 4; j <= idxNext; j++ {
			shift := j - 3
			n := b.whatsOn[j]
			b.whatsOn[shift] = n
			b.whereIs[n] = shift
		}

		idxp1 := idxNext - 2
		idxp2 := idxNext - 1
		idxp3 := idxNext
		b.whatsOn[idxp1] = p1
		b.whereIs[p1] = idxp1

		b.whatsOn[idxp2] = p2
		b.whereIs[p2] = idxp2

		b.whatsOn[idxp3] = p3
		b.whereIs[p3] = idxp3
	}

	// after having moved numbers, where's the current again?
	idxCurrentMoved := b.whereIs[b.current]

	b.current = b.whatsOn[((idxCurrentMoved)%b.length)+1]
	return b
}

func task2() {
	start := getInputs()
	bigOof := newBigOof(start, cupsInCircle)
	fmt.Printf("got our big oof\n")
	for i := 0; i < moveCupsThisManyTimes; i++ {

		bigOof = bigOof.step()
		fmt.Printf("simulated %d rounds so far...\n", i)

	}

	fmt.Printf("Day 23 Task 2: after the crab played ten million rounds of crab cups with one million cups, the product of the two cups clockwise to cup \"1\" is %d\n", gimmeProduct(bigOof))
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

func gimmeProduct(in bigOof) int {
	idxCurrent := in.whereIs[in.current]
	return in.whatsOn[(idxCurrent%in.length)+1] * in.whatsOn[((idxCurrent+1)%in.length)+1]
}
