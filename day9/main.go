package day9

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	filename       = "day9/input.txt"
	preambleLength = 25
)

func Tasks() {
	task1()

	task2()
}
func task1() {
	//inputs := getInputs()
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

func task2() {}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// strip the trailing newline
	sData := strings.TrimRight(string(data), "\n")
	stringSlice := strings.Split(sData, "\n")
	values := make([]int, 0, len(stringSlice))
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
