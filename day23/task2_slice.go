package day23

import "fmt"

func task2slice() {
	product := getProduct(createInputSlice(getInputs(), cupsInCircle), moveCupsThisManyTimes)
	fmt.Printf("day23 task 2: the product of the two cups clockwise to 1 is %d\n", product)
}

func createInputSlice(start []int, total int) []int {
	if total < len(start) {
		panic(fmt.Sprintf("total %d is smaller than the original slice length %d", total, len(start)))
	}

	for i := len(start) + 1; i <= total; i++ {
		start = append(start, i)
	}
	return start
}

func getProduct(in []int, n int) int {
	cg := newCupGame(in)
	for i := 0; i < n; i++ {
		cg.play()
	}

	p1 := cg.Elements[1]
	p2 := cg.Elements[p1]
	return p1 * p2
}
