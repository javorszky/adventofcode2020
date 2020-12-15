package day15

import (
	"fmt"
	"time"
)

const task2NthNumber = 30000000

func task2() {
	t := time.Now()
	num := nthNumber(getInputs(), task2NthNumber)
	end := time.Since(t)

	fmt.Printf("Day 15 task 2: last number is %d, took %s\n", num, end)
}
