package main

import (
	"io/ioutil"
	"strings"
)

const filename = "input.txt"

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

// isTree returns whether the ground at the given coordinates is a tree or not. Column is arbitrary number, is not
// constrained by width of the map.
func isTree(row, column int, slope []string) bool {
	return false
}
