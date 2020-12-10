package day10

import (
	"io/ioutil"
	"strings"
)

const filename = "day10/input.txt"

func Tasks() {
	task1()
	task2()
}

func task1() {
	_ = getInputs()
}

func task2() {
	_ = getInputs()
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
