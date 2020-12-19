package day19

import (
	"io/ioutil"
	"strings"
)

const filename = "day19/input.txt"

// getInputs reads the input.txt file and returns the rules and messages as slices of strings each.
func getInputs() ([]string, []string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	input := strings.TrimRight(string(data), "\n")
	parts := strings.Split(input, "\n\n")
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}
