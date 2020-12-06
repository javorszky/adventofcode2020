package day6

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "day6/input.txt"

func Tasks() {
	task1()
}

func task1() {
	count := 0
	for _, group := range getInputs() {
		yes := make(map[string]struct{})

		for _, person := range strings.Fields(group) {
			for _, answer := range strings.Split(person, "") {
				yes[answer] = struct{}{}
			}
		}

		count = count + len(yes)

	}
	fmt.Printf("day 6 task 1: total number of yes answers: %d\n\n", count)
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// strip the trailing newline
	sData := strings.TrimRight(string(data), "\n")

	return strings.Split(sData, "\n\n")
}
