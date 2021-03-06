package day17

import (
	"io/ioutil"
	"strings"
)

const (
	filename = "day17/input.txt"
	active   = "#"
	inactive = "."

	zStart = 0
	cycles = 6
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
