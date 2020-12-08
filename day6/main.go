package day6

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "day6/input.txt"

func Tasks() {
	task1()

	task2()
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
	fmt.Printf("\nDay 6 task 1: total number of yes answers: %d\n", count)
}

func task2() {
	count := 0
	for _, group := range getInputs() {
		yes := make(map[string]int)
		people := strings.Fields(group)
		lenp := len(people)
		for _, person := range people {
			for _, answer := range strings.Split(person, "") {
				yes[answer] = yes[answer] + 1
			}
		}

		for _, v := range yes {
			if v == lenp {
				count++
			}
		}
	}
	fmt.Printf("Day 6 task 2: number of answers where everyone answered yes to in a group: %d\n", count)
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
