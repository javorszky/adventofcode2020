package day15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day15/input.txt"

func Tasks() {
	task1()
	task2()
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	out := make([]int, 0)
	for _, n := range strings.Split(strings.TrimRight(string(data), "\n"), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(fmt.Sprintf("day 15: converting string %s to int failed: %s", n, err))
		}
		out = append(out, i)
	}

	return out
}
