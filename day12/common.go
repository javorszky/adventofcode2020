package day12

import (
	"io/ioutil"
	"strings"
)

func manhattanDistance(s ship) int {
	return abs(s.X) + abs(s.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
