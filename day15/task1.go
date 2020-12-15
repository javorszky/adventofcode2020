package day15

import "fmt"

const task1NthNumber = 2020

func task1() {
	num := nthNumber(getInputs(), task1NthNumber)

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
	for i := len(in) + 1; i < n; i++ {
		idx, ok := in[last]
		if !ok {
			in[last] = i
			last = 0
			continue
		}
		in[last] = i
		last = i - idx
	}

	return last
}
