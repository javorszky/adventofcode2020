package day9

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	filename       = "day9/input.txt"
	preambleLength = 25
)

func Tasks() {
	num, err := task1()
	if err != nil {
		fmt.Printf("\nDay 9 task 1: %s\n", err)
	}
	fmt.Printf("\nDay 9 task 1: the number %d is not valid\n", num)

	sum, err := task2(num)
	if err != nil {
		fmt.Printf("Day 9 task 2: %s\n", err)
	}
	fmt.Printf("Day 9 task 2: the sum of the lowest and highest number in the slice is %d\n", sum)
}

func task1() (int, error) {
	inputs := getInputs()

	for i, n := range inputs[preambleLength:] {
		valid := t1checkNumber(inputs[i:i+preambleLength], n)
		if !valid {
			return n, nil
		}
	}
	return 0, errors.New("no sum number not found")
}

func t1checkNumber(previous []int, num int) bool {
	// make a hash
	hash := make(map[int]struct{}, 0)
	for _, n := range previous {
		hash[n] = struct{}{}
	}

	for n := range hash {
		_, ok := hash[num-n]
		if ok && num-n != n {
			return true
		}
	}
	return false
}

func task2(num int) (int, error) {
	// find the contiguous list in the numbers and add the smallest and largest in the list.
	inputs := getInputs()
	for i := range inputs {
		sum := 0
		for j, n := range inputs[i:] {
			sum = sum + n
			if sum == num {
				cont := make([]int, len(inputs[i:i+j]))
				copy(cont, inputs[i:i+j])
				sort.Ints(cont)
				return cont[0] + cont[len(cont)-1], nil
			}
			if sum > num {
				break
			}
		}
	}
	return 0, errors.New("not found")
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// strip the trailing newline
	sData := strings.TrimRight(string(data), "\n")
	stringSlice := strings.Split(sData, "\n")
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
