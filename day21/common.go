package day21

import (
	"io/ioutil"
	"strings"
)

const filename = "day21/input.txt"

//const filename = "day21/input_example.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
