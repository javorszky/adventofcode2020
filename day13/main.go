package day13

import (
	"io/ioutil"
	"strings"
)

const filename = "day13/input.txt"

func Tasks() {
	task1()
	task2()
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
