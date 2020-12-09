package day9

import (
	"io/ioutil"
	"strings"
)

const filename = "day9/input.txt"

func Tasks() {
	task1()

	task2()
}
func task1() {}

func task2() {}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// strip the trailing newline
	sData := strings.TrimRight(string(data), "\n")

	return strings.Split(sData, "\n")
}
