package day15

import "fmt"

const task2NthNumber = 30000000

func task2() {
	num := nthNumber(getInputs(), task2NthNumber)

	fmt.Printf("Day 15 task 2: last number is %d\n", num)
}
