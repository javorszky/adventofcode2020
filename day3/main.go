package main

import (
	"io/ioutil"
	"strings"
)

const filename = "input.txt"
const startRow = 0
const stepRow = 1
const startCol = 0
const stepCol = 3

func main() {

}

// getInputs reads the input.txt file, and arranges the contents into a map of unique numbers.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func step(n int) (int, int) {
	return startRow + n*stepRow, startCol + n*stepCol
}

// isTree returns whether the ground at the given coordinates is a tree or not. Column is arbitrary number, is not
// constrained by width of the map.
func isTree(row, column int, slope []string) bool {
	if len(slope) < row {
		return false
	}
	trueColumn := column % len(slope[0])

	r := slope[row]

	return "#" == string(r[trueColumn])
}
