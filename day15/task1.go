package day15

import "fmt"

func task1() {

	n := 7
	num := nthNumber(getInputs(), n)

	fmt.Printf("\nDay 15 task 1: last number is %d\n", num)
}

func nthNumber(in []int, n int) int {
	last := 0
	s := make(map[int]int, 0)
	for i := 1; i <= len(in); i++ {
		s[last] = i - 1
		last = in[i-1]
		if i == n {
			return last
		}
	}

	return nthNewNumber(s, last, n)
}

func nthNewNumber(in map[int]int, last, n int) int {
	for i := len(in) + 1; i <= n; i++ {
		fmt.Printf("new: looking for %d, last number was %d\nmap: %#v\n", i, last, in)
		idx, ok := in[last]
		if !ok {
			fmt.Printf("new: last was not found, so setting 0 and last 0 as %d, and continuing\n", i)
			in[last] = i
			last = 0
			continue
		}
		in[last] = i
		last = i - idx
	}
	return last
}
