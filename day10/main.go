package day10

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "day10/input.txt"

func Tasks() {
	task1()
	task2()
}

func task1() {
	input := getInputs()

	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	ones, threes, previous := 0, 0, 0
	for _, k := range input {
		switch k - previous {
		case 1:
			ones++
		case 3:
			threes++
		}
		previous = k
	}

	fmt.Printf("\nDay 10 task 1: product of nDiff1 and nDiff3 is %d\n", ones*threes)
}

func task2() {
	input := getInputs()
	sort.Ints(input)
	device := input[len(input)-1] + 3
	input = append(input, device)
	input = append([]int{0}, input...)
	acc := countBranchEasier(input)
	fmt.Printf("Day 10 task 2: the total number of different combinations is %d\n", acc)
}

func countBranch(start, end int, joltageMap map[int]struct{}) int {
	var f func(int)
	branches := 0
	f = func(value int) {
		if value == end {
			branches++
		}
		if _, ok := joltageMap[value+1]; ok {
			f(value + 1)
		}
		if _, ok := joltageMap[value+2]; ok {
			f(value + 2)
		}
		if _, ok := joltageMap[value+3]; ok {
			f(value + 3)
		}
	}
	f(start)
	return branches
}

func countBranchEasier(input []int) int {
	fmt.Printf("%#v\n", input)
	joltageMapCount := make(map[int]int, 0)
	for _, j := range input {
		joltageMapCount[j] = 0
	}
	joltageMapCount[0] = 1
	end := input[len(input)-1]
	for _, v := range input[1:] {
		local := 0
		for i := 1; i < 4; i++ {
			c, ok := joltageMapCount[v-i]
			if ok {
				local = local + c
			}
		}
		joltageMapCount[v] = local
	}

	return joltageMapCount[end]
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	stringSlice := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	values := make([]int, len(stringSlice), len(stringSlice))

	for i, v := range stringSlice {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("could not turn string %s into an int: %s\n", v, err)
			os.Exit(1)
		}
		values[i] = num
	}
	return values
}
