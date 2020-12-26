package day23

import "fmt"

func task2slice() {

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
