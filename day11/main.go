package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "day11/input.txt"

type floorIsLave struct {
	Message string
}

// Error returns the string representation of the error.
func (f floorIsLave) Error() string {
	return f.Message
}

type outOfBounds struct {
	Message string
}

// Error returns the string representation of the error.
func (o outOfBounds) Error() string {
	return o.Message
}

func Tasks() {
	task1()
	task2()
}

func task1() {
	lineLength, input := t1normalize(getInputs())
	previous := input
	newState := ""
	for {
		newState = t1flip(lineLength, previous)

		if previous == newState {
			break
		}
		previous = newState
	}

	fmt.Printf("\nDay 11 task 1: once stabilized, there are %d occupied seats\n", strings.Count(newState, "#"))
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
